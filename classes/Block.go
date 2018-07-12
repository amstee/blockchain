package classes

import (
	"strconv"
	"crypto/sha256"
	"time"
	"bytes"
)

type Block struct {
	Timestamp int64
	Data []byte
	Hash [32]byte
	Prev *Block
	Iterations int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var headers []byte
	if b.Prev == nil {
		headers = bytes.Join([][]byte{{}, b.Data, timestamp}, []byte(""))

	} else {
		headers = bytes.Join([][]byte{b.Prev.Hash[:], b.Data, timestamp}, []byte(""))
	}
	hash := sha256.Sum256(headers)
	b.Hash = hash
}

func NewBlock(data string, prev *Block) *Block {
	block := &Block{time.Now().Unix(), []byte(data), [32]byte{}, prev, 0}
	pow := NewProofOfWork(block)
	i, hash := pow.Run()
	block.Hash = hash
	block.Iterations = i
	return block
}

func (b *Block) GetPrevHash() [32]byte {
	if b.Prev == nil {
		return [32]byte{}
	}
	return b.Prev.Hash
}