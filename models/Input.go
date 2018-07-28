package models

type TXInput struct {
	TxID		string `gorm:"primary_key"`
	OtxID		string
	Vout		int
	ScriptSig	string
}

func (itx *TXInput) GetTXID() []byte {
	return []byte(itx.TxID)
}

func (itx *TXInput) GetOTXID() []byte {
	return []byte(itx.OtxID)
}

func (itx *TXInput) CanUnlockOutput(data string) bool {
	return itx.ScriptSig == data
}