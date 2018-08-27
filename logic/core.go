package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
	"log"
	"github.com/amstee/blockchain/config"
)

func InitDatabase(db *gorm.DB) error {
	db.AutoMigrate(&models.BlockModel{})
	db.AutoMigrate(&models.TransactionModel{})
	db.AutoMigrate(&models.TXOutput{})
	db.AutoMigrate(&models.TXInput{})
	db.AutoMigrate(&models.Wallet{})
	return nil
}

func StartDatabase() *gorm.DB {
	db, err := gorm.Open(config.DbConf.BlocksDatabaseType, config.DbConf.BlocksDatabaseFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	InitDatabase(db)
	return db
}

func StartOutputsDatabase() *gorm.DB {
	db, err := gorm.Open(config.DbConf.OutputsDatabaseType, config.DbConf.OutputsDatabaseFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.AutoMigrate(&models.TXOutput{})
	return db
}

