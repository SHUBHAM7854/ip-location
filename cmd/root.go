package cmd

import (

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IPtracker CLI APP",
		Long: `IPtracker CLI APP.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}