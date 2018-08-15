package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
)

func DisplayWallets(db *gorm.DB) {
	wallets := models.GetWallets(db)
	wallets.Display()
}