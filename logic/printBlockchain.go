package logic

import (
	"github.com/amstee/blockchain/classes"
	"github.com/jinzhu/gorm"
)

func PrintBlockchain(db *gorm.DB, odb *gorm.DB) error {
	blockchain := classes.GetBlockChainFromGenesis(db, odb)
	blockchain.DisplayBlockChain()
	return nil
}