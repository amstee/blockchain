package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/logic"
)

var displayWalletsCmd = &cobra.Command{
	Use:   "display_wallets",
	Short: "Display all the wallets",
	Long:  `Display all the wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		db := logic.StartDatabase()
		defer db.Close()
		logic.DisplayWallets(db)
	},
}
