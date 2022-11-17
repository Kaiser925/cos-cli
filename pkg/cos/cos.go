package cos

import (
	"context"
	"io/fs"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/Kaiser925/cos-cli/pkg/trie"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const timeout = 10 * time.Second

type COS struct {
	cli   *cos.Client
	alias string
}

func NewCOS(alias string, URL string, secretID string, secretKey string) *COS {
	u, _ := url.Parse(URL)
	b := &cos.BaseURL{BucketURL: u}
	return &COS{
		cli: cos.NewClient(b, &http.Client{
			Timeout: timeout,
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,
				SecretKey: secretKey,
			},
		}),
		alias: alias,
	}
}

func (c *COS) ReadDir(ctx context.Context, name string) ([]fs.DirEntry, error) {
	b, _, err := c.cli.Bucket.Get(ctx, &cos.BucketGetOptions{})
	if err != nil {
		return nil, err
	}
	tree := trie.New[*objectFile](trie.PathSegment)
	for _, obj := range b.Contents {
		tree.Put(path.Join(c.alias, obj.Key), &objectFile{obj: obj})
	}

	t, ok := tree.Get(strings.TrimSuffix(name, "/"))
	if !ok {
		return nil, fs.ErrNotExist
	}

	if t.Value != nil && !t.Value.IsDir() {
		return []fs.DirEntry{t.Value}, nil
	}

	var dirs []fs.DirEntry
	for _, v := range t.Children {
		if v.Value != nil {
			dirs = append(dirs, v.Value)
		}
	}
	return dirs, nil
}

type objectFile struct {
	obj cos.Object
}

// implements fs.File for objectFile

func (o *objectFile) Read([]byte) (int, error) {
	// TODO implement me
	panic("implement me")
}

func (o *objectFile) Stat() (fs.FileInfo, error) { return o, nil }
func (o *objectFile) Close() error               { return nil }

// implements fs.FileInfo && fs.DirEntry for objectFile

func (o *objectFile) Type() fs.FileMode          { return o.Mode() }
func (o *objectFile) Info() (fs.FileInfo, error) { return o, nil }
func (o *objectFile) Name() string               { return path.Base(o.obj.Key) }
func (o *objectFile) Size() int64                { return o.obj.Size }
func (o *objectFile) IsDir() bool {
	return o.obj.Size == 0
}
func (o *objectFile) Sys() any { return nil }
func (o *objectFile) Mode() fs.FileMode {
	// Assume there is no write permission on objectFile
	if o.IsDir() {
		return fs.ModeDir | 0o444
	}
	return 0o444
}

func (o *objectFile) ModTime() time.Time {
	t, _ := time.Parse(time.RFC3339, o.obj.LastModified)
	return t
}

type fakeRootDir struct{}

func (o *fakeRootDir) Type() fs.FileMode  { return fs.ModeDir }
func (o *fakeRootDir) Name() string       { return "" }
func (o *fakeRootDir) Size() int64        { return 0 }
func (o *fakeRootDir) IsDir() bool        { return true }
func (o *fakeRootDir) Sys() any           { return nil }
func (o *fakeRootDir) Mode() fs.FileMode  { return fs.ModeDir }
func (o *fakeRootDir) ModTime() time.Time { return time.Now() }
