package classes

import (
	"math/big"
	"bytes"
	"github.com/amstee/blockchain/utils"
	"fmt"
	"crypto/sha256"
	"math"
)

const TargetBits = 16

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(i int) []byte {
	prevHash := pow.block.GetPrevHash()
	return bytes.Join([][]byte{
		prevHash[:],
		pow.block.Data,
		utils.IntegerToHex(pow.block.Timestamp),
		utils.IntegerToHex(int64(TargetBits)),
		utils.IntegerToHex(int64(i)),
	}, []byte(""))
}

func (pow *ProofOfWork) Run() (int, [32]byte) {
	var hashInt big.Int
	var hash [32]byte
	i := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for i < math.MaxInt64 {
		data := pow.prepareData(i)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x   %d", hash, i)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		i++
	}
	fmt.Print("\n\n")
	return i, hash
}

func (pow *ProofOfWork) Validate() bool {
	var hashint big.Int

	data := pow.prepareData(pow.block.Iterations)
	hash := sha256.Sum256(data)
	hashint.SetBytes(hash[:])
	isValid := hashint.Cmp(pow.target) == -1
	return isValid
}
