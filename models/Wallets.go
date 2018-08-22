package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Wallets struct {
	List map[string]*Wallet
}

func (ws *Wallets) Display() {
	i := 0

	for _, wallet := range ws.List {
		i += 1
		fmt.Printf("Wallet %d\n", i)
		wallet.Display()
	}
}

func GetWallets(db *gorm.DB) *Wallets {
	wallets := Wallets{}
	wallets.List = make(map[string]*Wallet)

	err := wallets.Load(db); if err != nil {
		return nil
	}
	return &wallets
}

func (ws *Wallets) GetWallet(address string) *Wallet {
	value, ok := ws.List[address]; if ok {
		return value
	}
	return nil
}

func (ws *Wallets) Save(db *gorm.DB) error {
	for _, wallet := range ws.List {
		if db.NewRecord(wallet) {
			err := db.Create(wallet).Error; if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ws *Wallets) Load(db *gorm.DB) error {
	var temp []*Wallet

	err := db.Find(&temp).Error; if err != nil {
		return err
	}
	for _, wallet := range temp {
		address := string(wallet.GetAddress())
		ws.List[address] = wallet
	}
	return nil
}
