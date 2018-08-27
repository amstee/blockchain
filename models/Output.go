package models

import (
	"github.com/jinzhu/gorm"
	"bytes"
	"github.com/amstee/blockchain/config"
	"github.com/amstee/blockchain/utils"
	"fmt"
)

type TXOutput struct {
	gorm.Model
	TxID			string
	Value 			int
	PubKeyHash		string
}

func (otx *TXOutput) GetKey() []byte {
	return []byte(otx.PubKeyHash)
}

func (otx *TXOutput) GetTXID() []byte {
	return []byte(otx.TxID)
}

func (otx *TXOutput) Lock(address []byte) {
	pubkeyhash := utils.Base58Decode(address)
	pubkeyhash = pubkeyhash[1 : len(pubkeyhash) - config.BlockchainConfig.CheckSumLen]
	otx.PubKeyHash = string(pubkeyhash)
}

func (otx *TXOutput) CanBeUnlocked(PubKeyHash []byte) bool {
	return bytes.Compare([]byte(otx.PubKeyHash), PubKeyHash) == 0
}

func (otx *TXOutput) Display() {
	fmt.Printf("Output TXID         : %x\n", otx.GetTXID())
	fmt.Printf("Output Value        : %d\n", otx.Value)
	fmt.Printf("Output PubKeyHashed : %x\n\n", otx.GetKey())
}