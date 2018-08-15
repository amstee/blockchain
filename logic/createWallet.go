package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
	"fmt"
)

func CreateWallet(db *gorm.DB) {
	wallet := models.NewWallet()
	if db.NewRecord(wallet) {
		err := db.Create(wallet).Error; if err != nil {
			fmt.Println(err)
		}
	}
	wallet.Display()
}
