package eventhandlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/m/v2/db"
	"github.com/m/v2/db/models"
	supply "github.com/m/v2/watchers/abi"
)

func TokensSupplyHandler(vLog types.Log) {
	supplyContractAbi, err := abi.JSON(strings.NewReader(supply.SupplyABI))
	if err != nil {
		fmt.Printf("Cannot get abi of contract: %s", err)
		os.Exit(1)
	}

	supplyDepositedEvent := supply.SupplyTokensSupplied{}
	supplyContractAbi.UnpackIntoInterface(&supplyDepositedEvent, "TokensSupplied", vLog.Data)

	initSupply := &models.Supply{
		Supplier:       supplyDepositedEvent.From.Hex(),
		TotalAmount:    supplyDepositedEvent.DepositInfo.TotalAmount.Uint64(),
		SupplyAmount:   supplyDepositedEvent.DepositInfo.SupplyAmount.Uint64(),
		LastSupplyTime: supplyDepositedEvent.DepositInfo.LastSupplyTime.Uint64(),
		Token:          supplyDepositedEvent.DepositInfo.Token.Hex(),
		MinAmount:      supplyDepositedEvent.DepositInfo.MinAmount.Uint64(),
		SplitTime:      supplyDepositedEvent.DepositInfo.SplitTime.Uint64(),
		Reciever:       supplyDepositedEvent.DepositInfo.Reciever.Hex(),
	}

	err = db.PostgresDBGlobalInstance.UpdateSupply(initSupply)
	if err != nil {
		fmt.Printf("Cannot insert supply: %s", err)
		os.Exit(1)
	}
}
