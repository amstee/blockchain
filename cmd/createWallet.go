package cmd

import (
"github.com/spf13/cobra"
"github.com/amstee/blockchain/logic"
)

var createWalletCmd = &cobra.Command{
	Use:   "create_wallet",
	Short: "Create new wallet",
	Long:  `Create new wallet`,
	Run: func(cmd *cobra.Command, args []string) {
		db := logic.StartDatabase()
		defer db.Close()
		logic.CreateWallet(db)
	},
}
