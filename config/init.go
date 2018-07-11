package config

import (
	"github.com/spf13/viper"
	"github.com/amstee/blockchain/utils"
)

func InitConfig(dbConf *DatabaseConfig) error {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.SetDefault("uri", "localhost")
	viper.SetDefault("port", 5000)
	viper.SetDefault("databaseType",  "sqlite3")
	viper.SetDefault("databaseFile", "sqlite.db")

	if err := viper.ReadInConfig(); err != nil {
		return utils.RaiseError("Unable to read in config file : ", err)
	}
	err := viper.Unmarshal(dbConf)
	if err != nil {
		return utils.RaiseError("Unable to decode config into struct : ", err)
	}
	return nil
}