package cmd

import (
	"github.com/spf13/cobra"
	"github.com/amstee/blockchain/logic"
)

var displayUnspentOutputsCmd = &cobra.Command{
	Use:   "display_unspent",
	Short: "display unspent outputs from outputs database",
	Long:  `display unspent outputs from outputs database`,
	Run: func(cmd *cobra.Command, args []string) {
		odb := logic.StartOutputsDatabase()
		defer odb.Close()
		logic.DisplayUnspentOutputs(odb)
	},
}
