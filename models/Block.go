package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"bytes"
	"crypto/sha256"
	"fmt"
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

func (b *BlockModel) Save(db *gorm.DB, odb *gorm.DB) {
	var temp TransactionModel
	var output TXOutput

	if db.NewRecord(b) {
		t := db.Create(b).Error; if t != nil {
			fmt.Println(t)
		}
	}
	for _, tx := range b.Transactions {
		// Check if newly spent outputs exist in odb and delete them
		if !tx.IsCoinbase() {
			for _, vin := range tx.Vin {
				err := db.Model(&vin).Related(&temp, "OtxID").Error; if err == nil {
					err = temp.LoadData(db); if err == nil {
						err = odb.First(&output, temp.Vout[vin.Vout].ID).Error; if err == nil {
							err = odb.Delete(&output).Error; if err != nil {
								fmt.Println("0", err)
							}
						}
					} else {
						fmt.Println("1", err)
					}
				} else {
					fmt.Println("2", err)
				}
			}
		}
		// Add created outputs to the odb
		var check TXOutput
		for _, out := range tx.Vout {
			if odb.First(&check, "id = ?",  out.ID).RecordNotFound() {
				odb.Create(&out)
			} else {
				fmt.Println("Record already exist in odb")
			}
		}
	}
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