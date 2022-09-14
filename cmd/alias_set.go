package cmd

import (
	"errors"
	"fmt"
	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
)

var aliasSet = &cobra.Command{
	Use:   "set",
	Short: "set a new alias to configuration file",
	Long: `set a new alias to configuration file

Example:
  cos-cli alias set mycos https://bucket.cos.region.myqcloud.com secret-id secret-key
`,
	PreRunE: preRunE,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 4 {
			return errors.New("incorrect number of arguments for alias set command")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name, uri, id, key := args[0], args[1], args[2], args[3]
		config.SetAlias(name, &config.AliasConfig{
			URL:       uri,
			SecretID:  id,
			SecretKey: key,
		})
		if err := config.Save(configFile); err != nil {
			fmt.Printf("save config failed: %v\n", err)
		}
		fmt.Printf("%s is ready\n", name)
	},
}
