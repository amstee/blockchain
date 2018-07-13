package classes

import (
	"fmt"
	"github.com/amstee/blockchain/models"
	"time"
	"github.com/jinzhu/gorm"
)

type Blockchain struct {
	blocks []*models.BlockModel
}

func (b *Blockchain) AddBlock(txs []*models.TransactionModel, db *gorm.DB) {
	prevBlock := b.blocks[len(b.blocks) - 1]
	newBlock := NewBlock(txs, prevBlock)
	b.blocks = append(b.blocks, newBlock)
	if db.NewRecord(newBlock) {
		db.Create(newBlock)
	}
}

func NewGenesisBlock(db *gorm.DB, address string) *models.BlockModel {
	var block models.BlockModel

	if err := db.Last(&block).Error; err != nil {
		coinbase := NewCoinBaseTX(address, "")
		block := NewBlock([]*models.TransactionModel{coinbase}, nil)
		if db.NewRecord(block) {
			db.Create(block)
		}
		return block
	}
	return &block
}

func GetBlockChainFromGenesis(db *gorm.DB) *Blockchain {
	var blocks []*models.BlockModel

	if err := db.Find(&blocks).Error; err != nil {
		return nil
	}
	return &Blockchain{blocks}
}

func NewBlockChain(db *gorm.DB, address string) *Blockchain {
	return &Blockchain{[]*models.BlockModel{NewGenesisBlock(db, address)}}
}

func (b *Blockchain) DisplayBlockChain(db *gorm.DB) {
	var txs []*models.TransactionModel
	var inputs []models.TXInput
	var outputs []models.TXOutput
	for _, block := range b.blocks {
		fmt.Printf("PrevHash: %x\n", block.GetPrevHash())
		fmt.Printf("Hash: %x\n", block.GetHash())
		db.Model(&block).Related(&txs, "BlockID")
		for _, tx := range txs {
			fmt.Printf("Transaction BLOCK   : %d\n", tx.BlockID)
			fmt.Printf("Transaction ID      : %x\n", tx.GetTXID())
			db.Model(&tx).Related(&inputs, "TxID")
			db.Model(&tx).Related(&outputs, "TxID")
			for _, itx := range inputs {
				fmt.Printf("Input TXID          : %x\n", itx.GetTXID())
				fmt.Printf("Input VOUT          : %x\n", itx.Vout)
				fmt.Printf("Input ScriptSig     : %s\n", itx.ScriptSig)
			}
			for _, otx := range outputs {
				fmt.Printf("Output TXID         : %x\n", otx.GetTXID())
				fmt.Printf("Output Value        : %x\n", otx.Value)
				fmt.Printf("Output ScriptPubKey : %s\n", otx.ScriptPubKey)
			}
		}
		fmt.Println()
	}
}

func NewBlock(txs []*models.TransactionModel, prev *models.BlockModel) *models.BlockModel {
	var id uint
	var prevHash string

	if prev == nil {
		id = 0
		prevHash = ""
	} else {
		id = prev.ID
		prevHash = prev.Hash
	}
	block := &models.BlockModel{
		Timestamp: time.Now().Unix(),
		Transactions: txs,
		Hash: "",
		Prev: id,
		PrevHash: prevHash,
		Iterations: 0,
	}
	pow := NewProofOfWork(block)
	i, hash := pow.Run()
	block.Hash = string(hash[:])
	block.Iterations = i
	return block
}