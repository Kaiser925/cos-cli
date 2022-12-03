package cmd

import (
	"fmt"

	"github.com/Kaiser925/cos-cli/pkg/output"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
)

var aliasLs = &cobra.Command{
	Use:     "ls",
	Short:   "list aliases",
	PreRunE: preRunE,
	Run: func(cmd *cobra.Command, args []string) {
		prefix := "  "
		fmt.Println("Config file:", configFile)
		for k, v := range config.Default().Aliases {
			fmt.Println(k + ":")
			kvs := []output.KVPair{
				{
					Key: "Bucket",
					Val: v.URL,
				},
				{
					Key: "SecretKey",
					Val: v.SecretKey,
				},
				{
					Key: "SecretID",
					Val: v.SecretID,
				},
			}
			output.KeyValues(prefix, kvs)
		}
	},
}
