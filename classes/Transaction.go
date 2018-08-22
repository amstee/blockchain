package classes

import (
	"github.com/amstee/blockchain/models"
	"fmt"
)

func NewCoinBaseTX(to, data string) *models.TransactionModel {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}
	txin := models.TXInput{TxID: "", OtxID: "", Vout: -1, Signature: "", PubKey: data}
	txout := models.TXOutput{TxID: "", Value: 50}
	txout.Lock([]byte(to))
	tx := models.TransactionModel{Txid: "", Vin: []models.TXInput{txin}, Vout: []models.TXOutput{txout}}
	tx.SetID()
	return &tx
}

func NewTransaction(from, to string, amount int, bc *Blockchain, ws *models.Wallets) *models.TransactionModel {
	var inputs []models.TXInput
	var outputs []models.TXOutput

	wallet := ws.GetWallet(from)
	PubKeyHash := wallet.GetPubKeyHashed()
	acc, validOutputs := bc.FindOutputs(PubKeyHash, amount); if acc < amount {
		return nil
	}
	for txid, outs := range validOutputs {
		for _, out := range outs {
			input := models.TXInput{TxID: "", OtxID: txid, Vout: out, Signature: "", PubKey: wallet.PublicKey}
			inputs = append(inputs, input)
		}
	}
	outputs = append(outputs, models.TXOutput{Value:amount, PubKeyHash: string(models.HashPubKey([]byte(to)))})
	if acc > amount {
		outputs = append(outputs, models.TXOutput{Value: acc - amount, PubKeyHash: string(PubKeyHash)})
	}
	tx := models.TransactionModel{Vin: inputs, Vout: outputs}
	tx.SetID()
	bc.SignTransaction(&tx, wallet.GetEcdsaPrivateKey())
	return &tx
}