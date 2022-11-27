package cmd

import (
	"github.com/spf13/cobra"
)

var alias = &cobra.Command{
	Use:   "alias",
	Short: "alias manage cos aliases",
}

func init() {
	alias.AddCommand(aliasLs)
	alias.AddCommand(aliasSet)
	alias.AddCommand(aliasRemove)
}
