package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
)

func AddBlocks(db *gorm.DB, args []string) {
	var blockchain *classes.Blockchain
	if blockchain = classes.GetBlockChainFromGenesis(db); blockchain == nil {
		blockchain = classes.NewBlockChain(db, "amstee")
	}
	blockchain.AddBlock(nil, db)
}
