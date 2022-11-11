package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/fatih/color"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/Kaiser925/cos-cli/pkg/cos"

	"github.com/spf13/cobra"
)

var ls = &cobra.Command{
	Use:     "ls",
	Short:   "list content in dir or bucket",
	PreRunE: preRunE,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) == 0 {
			name = "."
		} else {
			name = args[0]
		}
		list(context.Background(), name)
	},
}

func list(ctx context.Context, name string) {
	var entries []fs.DirEntry
	var err error

	ss := strings.Split(name, "/")
	if alias, ok := config.GetAlias(ss[0]); ok {
		entries, err = cos.NewCOS(alias.URL, alias.SecretID, alias.SecretKey).ReadDir(ctx, name)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		entries, err = os.ReadDir(name)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	t := table.NewWriter()
	style := table.StyleDefault
	style.Options.DrawBorder = false
	style.Box.MiddleVertical = ""
	t.SetStyle(style)

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	rows := make([]table.Row, 0, len(entries))
	for _, v := range entries {
		i, _ := v.Info()
		name := i.Name()
		modTime := i.ModTime().Format("[2006-01-02 15:04:05]")
		if i.IsDir() {
			name = color.CyanString(i.Name())
		}
		rows = append(rows, table.Row{modTime, i.Size(), name})
	}
	t.AppendRows(rows)
	fmt.Println(t.Render())
}
