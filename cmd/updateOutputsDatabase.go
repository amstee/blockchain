package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/logic"
)

var updateOutputsCmd = &cobra.Command{
	Use:   "update_outputs",
	Short: "update outputs database",
	Long:  `update outputs database`,
	Run: func(cmd *cobra.Command, args []string) {
		db := logic.StartDatabase()
		odb := logic.StartOutputsDatabase()
		defer db.Close()
		defer odb.Close()
		logic.UpdateOutputsDatabase(db, odb)
	},
}