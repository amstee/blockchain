package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type BlockModel struct {
	gorm.Model
	Timestamp 		int64
	Transactions 	[]*TransactionModel	`gorm:"foreignkey:BlockID"`
	Hash 			string				`gorm:"type:varchar(32)"`
	Prev 			uint
	PrevHash 		string 				`gorm:"type:varchar(32)"`
	Iterations 		int
}

func (b *BlockModel) LoadTransactions(db *gorm.DB) []*TransactionModel {
	var transactions []*TransactionModel

	db.Model(&b).Related(&transactions, "BlockID")
	for _, transaction := range transactions {
		db.Model(&transaction).Related(&transaction.Vin, "TxID")
		db.Model(&transaction).Related(&transaction.Vout, "TxID")
	}
	return transactions
}

func (b *BlockModel) GetPrevHash() []byte {
	return []byte(b.PrevHash)
}

func (b *BlockModel) GetHash() []byte {
	return []byte(b.Hash)
}

func (b *BlockModel) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, []byte(tx.Txid))
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *BlockModel) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var headers []byte
	if b.Prev == 0 {
		headers = bytes.Join([][]byte{{}, b.HashTransactions(), timestamp}, []byte(""))

	} else {
		headers = bytes.Join([][]byte{b.GetPrevHash(), b.HashTransactions(), timestamp}, []byte(""))
	}
	hash := sha256.Sum256(headers)
	b.Hash = string(hash[:])
}