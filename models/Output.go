package models

import "github.com/jinzhu/gorm"

type TXOutput struct {
	gorm.Model
	TxID			string
	Value 			int
	ScriptPubKey	string
}

func (otx *TXOutput) GetTXID() []byte {
	return []byte(otx.TxID)
}

func (otx *TXOutput) CanBeUnlocked(data string) bool {
	return otx.ScriptPubKey == data
}