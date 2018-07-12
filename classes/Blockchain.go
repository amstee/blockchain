package classes

import (
	"fmt"
	"strconv"
)

type Blockchain struct {
	blocks []*Block
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks) - 1]
	newBlock := NewBlock(data, prevBlock)
	b.blocks = append(b.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", nil)
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (b *Blockchain) DisplayBlockChain() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.GetPrevHash())
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}