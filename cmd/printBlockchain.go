package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/logic"
	"github.com/amstee/blockchain/config"
)

var printCmd = &cobra.Command{
	Use:   "printchain",
	Short: "print the whole blockchain",
	Long:  `print the whole blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		db := config.StartDatabase()
		defer db.Close()
		logic.PrintBlockchain(db)
	},
}