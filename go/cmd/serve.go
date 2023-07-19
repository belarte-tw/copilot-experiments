package cmd

import (
	"github.com/belarte-tw/copilot-experiments/server"
	"github.com/spf13/cobra"
)

// cobra command to start the server
// add the command to the root command

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Start("127.0.0.1:8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
