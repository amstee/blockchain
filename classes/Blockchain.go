package classes

import (
	"fmt"
	"strconv"
	"github.com/amstee/blockchain/models"
	"time"
)

type Blockchain struct {
	blocks []*models.BlockModel
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks) - 1]
	newBlock := NewBlock(data, prevBlock)
	b.blocks = append(b.blocks, newBlock)
}

func NewGenesisBlock() *models.BlockModel {
	return NewBlock("Genesis Block", nil)
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*models.BlockModel{NewGenesisBlock()}}
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