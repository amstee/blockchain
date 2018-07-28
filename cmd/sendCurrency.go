package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/config"
	"github.com/amstee/blockchain/logic"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send currency, Parans: 1 From, 2 to, 3 amount",
	Long:  `send currency`,
	Run: func(cmd *cobra.Command, args []string) {
		db := config.StartDatabase()
		defer db.Close()
		logic.SendCurrency(db, args)
	},
}