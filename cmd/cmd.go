package cmd

import (
	"os"
	"path"

	"github.com/Kaiser925/cos-cli/pkg/output"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
)

var (
	configDir  string
	configFile string
)

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
	rootCmd.AddCommand(ls)
}

// Execute executes cos-cli command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.Fatal(err.Error())
	}
}
