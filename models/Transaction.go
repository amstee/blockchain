package models

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

type TransactionModel struct {
	BlockID	uint		`gorm:"foreign_key"`
	Txid 	string		`gorm:"primary_key"`
	Vin		[]TXInput	`gorm:"foreignkey:TxID"`
	Vout	[]TXOutput	`gorm:"foreignkey:TxID"`
}

func (tx *TransactionModel) GetTXID() []byte {
	return []byte(tx.Txid)
}

func (tx *TransactionModel) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1
}

func (tx *TransactionModel) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.Txid = string(hash[:])
}