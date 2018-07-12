package logic

import (
	"github.com/amstee/blockchain/classes"
	"github.com/jinzhu/gorm"
)

func Run(db *gorm.DB) error {
	blockchain := classes.NewBlockChain()

	blockchain.AddBlock("Send 1 BTC to Ragnar")
	blockchain.AddBlock("Send 2 BTC to Ivar")

	blockchain.DisplayBlockChain()
	return nil
}