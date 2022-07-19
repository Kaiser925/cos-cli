package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cos-cli",
	Short: "cos-cli - Client for tencent cloud object storage and filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO(Kaiser925): subcommand add.
		fmt.Println("cos-cli is working in progress")
	},
}

// Execute executes cos-cli command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
