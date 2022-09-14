package cmd

import (
	"fmt"
	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var configDir string
var configFile string

var preRunE = func(cmd *cobra.Command, args []string) error {
	configFile = path.Join(configDir, "config.json")
	return config.LoadOrInit(configFile)
}

var rootCmd = &cobra.Command{
	Use:   "cos-cli",
	Short: "cos-cli - Client for tencent cloud object storage and filesystem",
}

func init() {
	home, _ := os.UserHomeDir()
	defaultDir := path.Join(home, ".coscli")
	rootCmd.PersistentFlags().StringVarP(&configDir, "config-dir", "C", defaultDir, "path to configuration folder")
	rootCmd.AddCommand(alias)
}

// Execute executes cos-cli command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
