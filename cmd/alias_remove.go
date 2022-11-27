package cmd

import (
	"fmt"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
)

var aliasRemove = &cobra.Command{
	Use:   "remove",
	Short: "remove alias from configuration file",
	Long: `remove alias from configuration file

Example:
  cos-cli alias remove mycos 
`,
	PreRunE: preRunE,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.RemoveAlias(args[0])
		if err := config.Save(configFile); err != nil {
			fmt.Printf("save config failed: %v\n", err)
		}
	},
}
