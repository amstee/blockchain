package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"github.com/amstee/blockchain/config"
	"github.com/jinzhu/gorm"
	"github.com/amstee/ecdsa-serializer"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"github.com/amstee/blockchain/utils"
)

type Wallet struct {
	gorm.Model
	PrivateKey string
	PublicKey string
}

func (w *Wallet) Display() {
	fmt.Printf("Wallet PubKeyHahed   : %x\n", w.GetPubKeyHashed())
	fmt.Printf("Wallet Address       : %s\n", w.GetAddress())
	fmt.Printf("Wallet PublicKey     : %x\n\n", w.GetKey())
}

func (w *Wallet) GetPubKeyHashed() []byte {
	return HashPubKey([]byte(w.PublicKey))
}

func (w *Wallet) GetKey() []byte {
	return []byte(w.PublicKey)
}

func (w *Wallet) GetEcdsaPrivateKey() *ecdsa.PrivateKey {
	pk, _ := EcdsaEncoder.DecodePrivKey(w.PrivateKey)

	return pk
}

func (w Wallet) GetAddress() []byte {
	version := []byte{config.BlockchainConfig.Version}

	hash := HashPubKey([]byte(w.PublicKey))

	vers := append(version, hash...)
	checksum := CheckSum(vers)

	full := append(vers, checksum...)
	address := utils.Base58Encode(full)
	return address
}

func HashPubKey(pubkey []byte) []byte {
	public := sha256.Sum256(pubkey)
	rip := ripemd160.New()
	_, err := rip.Write(public[:]); if err != nil {
		log.Fatalf("An error occured (RIPEMD160)")
	}
	prip := rip.Sum(nil)
	return prip
}

func CheckSum(payload []byte) []byte {
	f := sha256.Sum256(payload)
	s := sha256.Sum256(f[:])
	return s[:config.BlockchainConfig.CheckSumLen]
}

func NewWallet() *Wallet {
	private, public := NewKeyPair()
	privstr, _ := EcdsaEncoder.EncodePrivKey(private)
	wallet := Wallet{PrivateKey: privstr, PublicKey: string(public)}
	return &wallet
}

func NewKeyPair() (*ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader); if err != nil {
		log.Fatalf("An error occured generating the key pair")
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return private, pubKey
}