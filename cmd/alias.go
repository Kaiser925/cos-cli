package cmd

import (
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/spf13/cobra"
)

var alias = &cobra.Command{
	Use:   "alias",
	Short: "alias manage cos aliases",
}

var aliasLs = &cobra.Command{
	Use:   "ls",
	Short: "list aliases",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadOrInit(config.DefaultConfigFile())
		if err != nil {
			log.Fatal(err.Error())
		}
		prefix := "  "
		maxLen := 12
		for k, v := range cfg.Aliases {
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

func initAlias() {
	alias.AddCommand(aliasLs)
}
