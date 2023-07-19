package cmd

import "github.com/spf13/cobra"

// command line application using cobra
// use "copilot" as the command name
var rootCmd = &cobra.Command{
	Use:   "copilot",
	Short: "Copilot is a command line application",
}

// Execute executes the root command.
// The function does not return any value.
// Check for errors executing the command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
