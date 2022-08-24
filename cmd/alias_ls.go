package cmd

import (
	"fmt"
	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var aliasLs = &cobra.Command{
	Use:     "ls",
	Short:   "list aliases",
	PreRunE: preRunE,
	Run: func(cmd *cobra.Command, args []string) {
		prefix := "  "
		maxLen := 12
		for k, v := range config.Default().Aliases {
			fmt.Println(k + ":")
			printLine(prefix, "Bucket", v.BucketName, maxLen)
			printLine(prefix, "Region", v.Region, maxLen)
			printLine(prefix, "SecretKey", v.SecretKey, maxLen)
			printLine(prefix, "SecretID", v.SecretID, maxLen)
		}
	},
}

func printLine(prefix, key, val string, maxLength int) {
	fmt.Println(prefix + text.AlignJustify.Apply(fmt.Sprintf("%s :", key), maxLength) + " " + val)
}
