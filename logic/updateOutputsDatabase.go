package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/classes"
)

func UpdateOutputsDatabase(db *gorm.DB, odb *gorm.DB) {
	blockchain := classes.GetBlockChainFromGenesis(db, odb)
	blockchain.UpdateBlocks()
}