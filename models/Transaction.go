package models

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"crypto/elliptic"
	"github.com/jinzhu/gorm"
)

type TransactionModel struct {
	BlockID	uint		`gorm:"foreign_key"`
	Txid 	string		`gorm:"primary_key"`
	Vin		[]TXInput	`gorm:"foreignkey:TxID"`
	Vout	[]TXOutput	`gorm:"foreignkey:TxID"`
}

func (tx *TransactionModel) LoadData(db *gorm.DB) error {
	err := db.Model(tx).Related(&tx.Vin, "TxID").Error; if err != nil {
		return err
	}
	err2 := db.Model(tx).Related(&tx.Vout, "TxID").Error; if err2 != nil {
		return err2
	}
	return nil
}

func (tx *TransactionModel) GetTXID() []byte {
	return []byte(tx.Txid)
}

func (tx *TransactionModel) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1
}

func (tx *TransactionModel) Verify(prevTransactions map[string]*TransactionModel) bool {
	if !tx.IsCoinbase() {
		txCopy := tx.GetTrimmedTransaction()
		curve := elliptic.P256()

		for i, vin := range tx.Vin {
			prevTransaction := prevTransactions[vin.OtxID]
			txCopy.Vin[i].PubKey = prevTransaction.Vout[txCopy.Vin[i].Vout].PubKeyHash
			txCopy.SetID()
			txCopy.Vin[i].PubKey = ""

			r := big.Int{}
			s := big.Int{}
			sigLen := len(vin.Signature)
			r.SetBytes([]byte(vin.Signature)[:(sigLen / 2)])
			s.SetBytes([]byte(vin.Signature)[(sigLen / 2):])

			x := big.Int{}
			y := big.Int{}
			keyLen := len(vin.PubKey)
			x.SetBytes([]byte(vin.PubKey)[:(keyLen / 2)])
			y.SetBytes([]byte(vin.PubKey)[(keyLen / 2):])
			pubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}

			if ecdsa.Verify(&pubKey, []byte(txCopy.Txid), &r, &s) == false {
				return false
			}
		}
	}
	return true
}

func (tx *TransactionModel) GetTrimmedTransaction() *TransactionModel {
	var inputs []TXInput
	var outputs []TXOutput

	for _, input := range tx.Vin {
		inputs = append(inputs, TXInput{TxID: "", OtxID: input.OtxID, Vout: input.Vout, Signature: "", PubKey: ""})
	}
	for _, output := range tx.Vout {
		outputs = append(outputs, TXOutput{Value: output.Value, PubKeyHash: output.PubKeyHash})
	}

	return &TransactionModel{Txid: tx.Txid, Vin: inputs, Vout: outputs}
}

func (tx *TransactionModel) Sign(privKey *ecdsa.PrivateKey, prevTransactions map[string]*TransactionModel) {
	if !tx.IsCoinbase() {
		TXCopy := tx.GetTrimmedTransaction()

		for i := 0; i < len(TXCopy.Vin); i++ {
			prevTransaction := prevTransactions[TXCopy.Vin[i].OtxID]
			if TXCopy.Vin[i].Vout >= len(prevTransaction.Vout) {
				log.Fatalf("Invalid number of outputs in transaction, cannot sign")
			}
			TXCopy.Vin[i].PubKey = prevTransaction.Vout[TXCopy.Vin[i].Vout].PubKeyHash
			TXCopy.SetID()
			TXCopy.Vin[i].PubKey = ""

			r, s, err := ecdsa.Sign(rand.Reader, privKey, []byte(TXCopy.Txid)); if err != nil {
				log.Fatalf("An error occured with ecdsa signing %s", err)
			}
			signature := append(r.Bytes(), s.Bytes()...)
			tx.Vin[i].Signature = string(signature)
		}
	}
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