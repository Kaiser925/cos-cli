package cmd

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/Kaiser925/cos-cli/pkg/config"
	"github.com/Kaiser925/cos-cli/pkg/oss"

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
	var fileSystem fs.FS
	if alias, ok := config.GetAlias(name); ok {
		fileSystem = oss.NewCOS(alias.URL, alias.SecretID, alias.SecretKey)
	} else {
		fileSystem = os.DirFS(name)
	}

	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(path, ".") {
			return nil
		}

		log.Println(path)
		// not walk subdir
		if d.IsDir() {
			return fs.SkipDir
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getBucket(name string) string {
	name = path.Clean(name)
	return name
}
