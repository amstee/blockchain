package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
)

func CreateBlockchain(db *gorm.DB, odb *gorm.DB, args []string) {
	var blockchain *classes.Blockchain

	if blockchain = classes.GetBlockChainFromGenesis(db, odb); blockchain == nil {
		blockchain = classes.NewBlockChain(db, odb, args[0])
	}
}
