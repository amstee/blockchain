package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/config"
	"github.com/amstee/blockchain/logic"
)

var createBlockchainCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new blocks to the blockchain",
	Long:  `Create new blocks to the blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		db := config.StartDatabase()
		defer db.Close()
		logic.CreateBlockchain(db, args)
	},
}