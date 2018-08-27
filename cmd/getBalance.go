package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/logic"
)

var getBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Display balance for a given address",
	Long:  `Display balance for a given address`,
	Run: func(cmd *cobra.Command, args []string) {
		db := logic.StartDatabase()
		odb := logic.StartOutputsDatabase()
		defer db.Close()
		defer odb.Close()
		logic.GetBalance(db, odb, args)
	},
}