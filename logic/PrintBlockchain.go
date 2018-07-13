package logic

import (
	"github.com/amstee/blockchain/classes"
	"github.com/jinzhu/gorm"
)

func PrintBlockchain(db *gorm.DB) error {
	blockchain := classes.GetBlockChainFromGenesis(db)
	blockchain.DisplayBlockChain()
	return nil
}