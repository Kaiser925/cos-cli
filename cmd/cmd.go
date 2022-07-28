package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cos-cli",
	Short: "cos-cli - Client for tencent cloud object storage and filesystem",
}

func init() {
	initAlias()
	rootCmd.AddCommand(alias)
}

// Execute executes cos-cli command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
