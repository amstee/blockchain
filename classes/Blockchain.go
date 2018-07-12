package classes

import (
	"fmt"
	"strconv"
	"github.com/amstee/blockchain/models"
	"time"
	"github.com/jinzhu/gorm"
)

type Blockchain struct {
	blocks []*models.BlockModel
}

func (b *Blockchain) AddBlock(data string, db *gorm.DB) {
	prevBlock := b.blocks[len(b.blocks) - 1]
	newBlock := NewBlock(data, prevBlock)
	b.blocks = append(b.blocks, newBlock)
	if db.NewRecord(newBlock) {
		db.Create(newBlock)
	}
}

func NewGenesisBlock(db *gorm.DB) *models.BlockModel {
	var block models.BlockModel

	if err := db.Last(&block).Error; err != nil {
		block := NewBlock("Genesis Block", nil)
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
		return NewBlockChain(db)
	}
	return &Blockchain{blocks}
}

func NewBlockChain(db *gorm.DB) *Blockchain {
	return &Blockchain{[]*models.BlockModel{NewGenesisBlock(db)}}
}

func (b *Blockchain) DisplayBlockChain() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.GetPrevHash())
		fmt.Printf("Hash: %x\n", block.GetHash())
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func NewBlock(data string, prev *models.BlockModel) *models.BlockModel {
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
		Data: data,
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