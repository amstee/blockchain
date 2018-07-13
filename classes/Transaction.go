package classes

import (
	"github.com/amstee/blockchain/models"
	"fmt"
)

func NewCoinBaseTX(to, data string) *models.TransactionModel {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}
	txin := models.TXInput{TxID: "", Vout: -1, ScriptSig: data}
	txout := models.TXOutput{Value: 50, ScriptPubKey: to}
	tx := models.TransactionModel{Txid: "", Vin: []models.TXInput{txin}, Vout: []models.TXOutput{txout}}
	tx.SetID()
	return &tx
}
