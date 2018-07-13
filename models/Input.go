package models

import "github.com/jinzhu/gorm"

type TXInput struct {
	gorm.Model
	TxID		string
	Vout		int
	ScriptSig	string
}

func (itx *TXInput) GetTXID() []byte {
	return []byte(itx.TxID)
}

func (itx *TXInput) CanUnlockOutput(data string) bool {
	return itx.ScriptSig == data
}