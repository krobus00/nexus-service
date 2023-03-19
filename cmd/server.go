package cmd

import (
	"github.com/krobus00/nexus-service/internal/bootstrap"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "nexus graphql server",
	Long:  `nexus graphql server`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
