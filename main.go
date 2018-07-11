package main

import (
	"github.com/amstee/blockchain/config"
	"log"
)

func main() {
	dbConf := new(config.DatabaseConfig)
	err := config.InitConfig(dbConf); if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Database connection uri is %s", dbConf.Uri)
	log.Printf("Database connection port is %d", dbConf.Port)
	log.Printf("DatabaseType : %s", dbConf.DatabaseType)
	log.Printf("DatabaseFile : %s", dbConf.DatabaseFile)
}
