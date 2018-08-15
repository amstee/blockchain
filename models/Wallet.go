package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"crypto"
	"github.com/amstee/blockchain/config"
	"github.com/itchyny/base58-go"
	"github.com/jinzhu/gorm"
	"github.com/amstee/ecdsa-serializer"
	"fmt"
)

type Wallet struct {
	gorm.Model
	PrivateKey string
	PublicKey string
}

func (w *Wallet) Display() {
	fmt.Printf("Input PrivateKey    : %s\n", w.PrivateKey)
	fmt.Printf("Input PublicKey     : %x\n", w.GetPubKeyHashed())
	fmt.Printf("Input Address       : %x\n", w.GetAddress())
}

func (w *Wallet) GetPubKeyHashed() []byte {
	return HashPubKey([]byte(w.PublicKey))
}

func (w *Wallet) GetEcdsaPrivateKey() *ecdsa.PrivateKey {
	pk, _ := EcdsaEncoder.DecodePrivKey(w.PrivateKey)

	return pk
}

func (w Wallet) GetAddress() []byte {
	version := []byte(config.BlockchainConfig.Version)
	hash := HashPubKey([]byte(w.PublicKey))
	vers := append(version, hash...)
	checksum := CheckSum(vers)

	full := append(vers, checksum...)
	encoder := base58.BitcoinEncoding
	address, err := encoder.Encode(full); if err != nil {
		log.Fatalf("Error encoding in base58")
	}
	return address
}

func HashPubKey(pubkey []byte) []byte {
	public := sha256.Sum256(pubkey)
	rip := crypto.RIPEMD160.New()
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
	privstr, _ := EcdsaEncoder.EncodePrivKey(&private)
	wallet := Wallet{PrivateKey: privstr, PublicKey: string(public)}
	return &wallet
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader); if err != nil {
		log.Fatalf("An error occured generating the key pair")
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}