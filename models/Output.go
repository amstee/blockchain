package models

import (
	"github.com/jinzhu/gorm"
	"bytes"
	"github.com/itchyny/base58-go"
	"log"
	"github.com/amstee/blockchain/config"
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
	encoder := base58.BitcoinEncoding

	pubkeyhash, err := encoder.Decode(address); if err != nil {
		log.Fatalf("Error decoding from base58")
	}
	pubkeyhash = pubkeyhash[1 : len(pubkeyhash) - config.BlockchainConfig.CheckSumLen]
	otx.PubKeyHash = string(pubkeyhash)
}

func (otx *TXOutput) CanBeUnlocked(PubKeyHash []byte) bool {
	return bytes.Compare([]byte(otx.PubKeyHash), PubKeyHash) == 0
}