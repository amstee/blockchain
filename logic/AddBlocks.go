package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
)

func AddBlocks(db *gorm.DB, args []string) {
	blockchain := classes.NewBlockChain(db)
	for _, s := range args {
		blockchain.AddBlock(s, db)
	}
}
