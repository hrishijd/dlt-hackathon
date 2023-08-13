package supplybot

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m/v2/config"
	"github.com/m/v2/db"
	supply_contract "github.com/m/v2/watchers/abi"
)

func InitializeSupplyBot(wg *sync.WaitGroup) {
	config := config.ReadBotConfig()

	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		fmt.Printf("Cannot connect to RPC: %s", err)
		os.Exit(1)
	}

	Run(client, config)

	wg.Done()
}

func Run(client *ethclient.Client, botConfig config.BotConfig) {
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Cannot get chain id: %s", err)
		os.Exit(1)
	}

	privateKey, err := crypto.HexToECDSA(botConfig.PrivateKeyOwner)
	if err != nil {
		fmt.Printf("Cannot get private key: %s", err)
		os.Exit(1)
	}

	supplyContract, err := supply_contract.NewSupply(common.HexToAddress(botConfig.Contract), client)
	if err != nil {
		fmt.Printf("Cannot connect to contract: %s", err)
		os.Exit(1)
	}

	fmt.Println("Supply Bot is Running")
	for {
		supplies, err := db.PostgresDBGlobalInstance.GetSupplysSortedByPriority()
		if err != nil {
			fmt.Printf("Cannot fetch supplies: %s", err)
			os.Exit(1)
		}
		for _, supply := range supplies {
			if !supply.TxSubmitted {
				_, err := supplyContract.SupplyTokens(&bind.TransactOpts{
					From: common.HexToAddress(botConfig.PublicKeyOwner),
					Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
						return types.SignTx(tx, types.LatestSignerForChainID(chainId), privateKey)
					},
				}, common.HexToAddress(supply.Supplier), big.NewInt(time.Now().Unix()))
				if err != nil {
					fmt.Printf("Cannot initiate : %s", err)
					os.Exit(1)
				}
			}
		}
	}
}
