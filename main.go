package main

import (
	"github.com/amstee/blockchain/config"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/amstee/blockchain/logic"
)

func main() {
	dbConf := new(config.DatabaseConfig)
	err := config.InitConfig(dbConf); if err != nil {
		log.Fatalf(err.Error())
	}
	db, err := gorm.Open(dbConf.DatabaseType, dbConf.DatabaseFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()
	config.InitDatabase(db)
	logic.Run(db)
}