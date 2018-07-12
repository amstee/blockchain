package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/amstee/blockchain/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/amstee/blockchain/logic"
)

var CfgFile = ""

var rootCmd = &cobra.Command{
	Use: "blockchain",
	Short: "A CLI go blockchain implementation",
	Long: "A CLI go blockchain implementation",
	Run: func(cmd *cobra.Command, args []string) {
		db := config.StartDatabase()
		defer db.Close()
		config.InitDatabase(db)
		logic.PrintBlockchain(db)
	},
}

func initConfig() {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err := viper.Unmarshal(config.DbConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "config.json", "config file")
	viper.SetConfigFile(CfgFile)
	viper.AddConfigPath(".")
	viper.SetDefault("uri", "localhost")
	viper.SetDefault("port", 5000)
	viper.SetDefault("databaseType",  "sqlite3")
	viper.SetDefault("databaseFile", "sqlite.db")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}