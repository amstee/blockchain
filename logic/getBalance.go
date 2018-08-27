package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
	"fmt"
	"github.com/amstee/blockchain/utils"
)

func GetBalance(db *gorm.DB, odb *gorm.DB, args []string) error {
	blockchain := classes.GetBlockChainFromGenesis(db, odb)
	blockchain.SetOutputsDB(odb)
	balance := 0

	pubKeyHash := utils.Base58Decode([]byte(args[0]))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	outputs := blockchain.GetUnspentOutputs(pubKeyHash)
	for _, out := range outputs {
		balance += out.Value
	}
	fmt.Printf("Balance for address %s is %d\n", args[0], balance)
	return nil
}