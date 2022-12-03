package cmd

import (
	"context"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/Kaiser925/cos-cli/pkg/output"

	"github.com/Kaiser925/cos-cli/pkg/tool"

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
		entries, err = cos.NewCOS(ss[0], alias.URL, alias.SecretID, alias.SecretKey).ReadDir(ctx, name)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		info, err := os.Stat(name)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if !info.IsDir() {
			entries = []fs.DirEntry{fileEntry(info)}
		} else {
			entries, err = os.ReadDir(name)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	rows := make([][]any, 0, len(entries))
	for _, v := range entries {
		i, _ := v.Info()
		name := i.Name()
		size := tool.GetByteSize(i.Size())
		modTime := i.ModTime().Format("[2006-01-02 15:04:05]")
		if i.IsDir() {
			name = color.CyanString(i.Name())
		}
		rows = append(rows, table.Row{modTime, size, name})
	}
	output.Table(rows)
}

type wrapper struct {
	info fs.FileInfo
}

func (w wrapper) Name() string {
	return w.info.Name()
}

func (w wrapper) IsDir() bool {
	return w.info.IsDir()
}

func (w wrapper) Type() fs.FileMode {
	return w.info.Mode()
}

func (w wrapper) Info() (fs.FileInfo, error) {
	return w.info, nil
}

func fileEntry(info fs.FileInfo) fs.DirEntry {
	return wrapper{info: info}
}
