package models

import "bytes"

type TXInput struct {
	TxID		string `gorm:"primary_key"`
	OtxID		string
	Vout		int
	Signature	string
	PubKey 		string
}

func (itx *TXInput) GetTXID() []byte {
	return []byte(itx.TxID)
}

func (itx *TXInput) GetOTXID() []byte {
	return []byte(itx.OtxID)
}

func (itx *TXInput) CanUnlockOutput(hash []byte) bool {
	lock := HashPubKey([]byte(itx.PubKey))
	return bytes.Compare(lock, hash) == 0
}

func (itx *TXInput) GetKey() []byte {
	return []byte(itx.PubKey)
}

func (itx *TXInput) GetSignature() []byte {
	return []byte(itx.Signature)
}