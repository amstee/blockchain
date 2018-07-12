package logic

import (
	"github.com/amstee/blockchain/classes"
	"github.com/jinzhu/gorm"
)

func Run(db *gorm.DB) error {
	blockchain := classes.GetBlockChainFromGenesis(db)
	//blockchain := classes.NewBlockChain(db)

	//blockchain.AddBlock("Send 1 BTC to Ragnar", db)
	//blockchain.AddBlock("Send 2 BTC to Ivar", db)

	blockchain.DisplayBlockChain()
	return nil
}