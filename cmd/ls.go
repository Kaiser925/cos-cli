package cmd

import (
	"context"
	"io/fs"
	"log"
	"os"
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
	var fsys fs.FS

	if alias, ok := config.GetAlias(name); ok {
		var err error
		fsys, err = oss.NewCOS(alias.URL, alias.SecretID, alias.SecretKey).BucketFS(ctx, "oss")
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		fsys = os.DirFS(name)
	}

	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
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
