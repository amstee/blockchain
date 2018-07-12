package config

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
)

type DatabaseConfig struct {
	Port int
	Uri string
	DatabaseType string
	DatabaseFile string
}

func InitDatabase(db *gorm.DB) error {
	db.AutoMigrate(&models.BlockModel{})
	return nil
}