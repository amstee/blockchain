package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
	"fmt"
)

func GetBalance(db *gorm.DB, args []string) error {
	blockchain := classes.GetBlockChainFromGenesis(db)
	balance := 0
	outputs := blockchain.GetUnspentOutputs(args[0])

	for _, out := range outputs {
		balance += out.Value
	}
	fmt.Printf("Balance for address %s is %d\n", args[0], balance)
	return nil
}