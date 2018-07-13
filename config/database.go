package config

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
	"log"
)

type DatabaseConfig struct {
	Port int
	Uri string
	DatabaseType string
	DatabaseFile string
}

func InitDatabase(db *gorm.DB) error {
	db.AutoMigrate(&models.BlockModel{})
	db.AutoMigrate(&models.TransactionModel{})
	db.AutoMigrate(&models.TXOutput{})
	db.AutoMigrate(&models.TXInput{})
	return nil
}

func StartDatabase() *gorm.DB {
	db, err := gorm.Open(DbConf.DatabaseType, DbConf.DatabaseFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	InitDatabase(db)
	return db
}

var DbConf = new(DatabaseConfig)
