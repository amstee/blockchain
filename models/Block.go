package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type BlockModel struct {
	gorm.Model
	Timestamp int64
	Data string
	Hash string		`gorm:"type:varchar(32)"`
	Prev uint
	PrevHash string `gorm:"type:varchar(32)"`
	Iterations int
}

func (b *BlockModel) GetPrevHash() []byte {
	return []byte(b.PrevHash)
}

func (b *BlockModel) GetHash() []byte {
	return []byte(b.Hash)
}

func (b *BlockModel) GetData() []byte {
	return []byte(b.Data)
}

func (b *BlockModel) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var headers []byte
	if b.Prev == 0 {
		headers = bytes.Join([][]byte{{}, b.GetData(), timestamp}, []byte(""))

	} else {
		headers = bytes.Join([][]byte{b.GetPrevHash(), b.GetData(), timestamp}, []byte(""))
	}
	hash := sha256.Sum256(headers)
	b.Hash = string(hash[:])
}