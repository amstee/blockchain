package classes

import (
	"github.com/amstee/blockchain/models"
	"fmt"
)

func NewCoinBaseTX(to, data string) *models.TransactionModel {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}
	txin := models.TXInput{TxID: "", OtxID: "", Vout: -1, ScriptSig: data}
	txout := models.TXOutput{TxID: "", Value: 50, ScriptPubKey: to}
	tx := models.TransactionModel{Txid: "", Vin: []models.TXInput{txin}, Vout: []models.TXOutput{txout}}
	tx.SetID()
	return &tx
}

func NewTransaction(from, to string, amount int, bc *Blockchain) *models.TransactionModel {
	var inputs []models.TXInput
	var outputs []models.TXOutput

	acc, validOutputs := bc.FindOutputs(from, amount); if acc < amount {
		return nil
	}
	for txid, outs := range validOutputs {
		for _, out := range outs {
			input := models.TXInput{TxID: "", OtxID: txid, Vout: out, ScriptSig: from}
			inputs = append(inputs, input)
		}
	}
	outputs = append(outputs, models.TXOutput{Value:amount, ScriptPubKey: to})
	if acc > amount {
		outputs = append(outputs, models.TXOutput{Value: acc - amount, ScriptPubKey: from})
	}
	tx := models.TransactionModel{Vin: inputs, Vout: outputs}
	tx.SetID()
	return &tx
}