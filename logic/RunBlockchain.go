package logic

import (
	"github.com/amstee/blockchain/classes"
)

func Run() error {
	blockchain := classes.NewBlockChain()

	blockchain.AddBlock("Send 1 BTC to Ragnar")
	blockchain.AddBlock("Send 2 BTC to Ivar")

	blockchain.DisplayBlockChain()
	return nil
}