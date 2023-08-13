package watchers

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m/v2/config"
	eventhandlers "github.com/m/v2/event-handlers"
)

func InitializeWatcher(wg *sync.WaitGroup) {
	config := config.ReadWatcherConfig()

	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		fmt.Printf("Cannot connect to RPC: %s", err)
		os.Exit(1)
	}

	contractAddress := common.HexToAddress(config.Contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(config.BlockNumber),
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Printf("Cannot subscripbe to events: %s", err)
		os.Exit(1)
	}

	WatchEvents(sub, logs)

	wg.Done()
}

func WatchEvents(sub ethereum.Subscription, logs chan types.Log) {

	SupplyDepositedSig := "tokensDeposited(address,DepositInfo)"
	SupplyDepositedHash := crypto.Keccak256Hash([]byte(SupplyDepositedSig))

	SupplySuppliedSig := "tokensSupplied(address,DepositInfo)"
	SupplySuppliedSigHash := crypto.Keccak256Hash([]byte(SupplySuppliedSig))

	fmt.Println("Supply watcher is Running")

	for {
		select {
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case SupplyDepositedHash.Hex():
				fmt.Println("Recieved Opened supply")
				eventhandlers.TokenDepositHandler(vLog)
			case SupplySuppliedSigHash.Hex():
				fmt.Println("Recieved Joined supply")
				eventhandlers.TokensSupplyHandler(vLog)
			}
		}
	}
}
