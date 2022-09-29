package oss

import (
	"context"
	"io/fs"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const timeout = 10 * time.Second

type COS struct {
	cli *cos.Client
}

func NewCOS(URL string, secretID string, secretKey string) *COS {
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
	}
}

// BucketFS returns a file system (a fs.FS) for the tree of files rooted at the bucket.
func (c *COS) BucketFS(ctx context.Context, name string) (fs.StatFS, error) {
	opt := &cos.BucketGetOptions{}
	b, _, err := c.cli.Bucket.Get(ctx, opt)
	if err != nil {
		return nil, err
	}
	return &cosFS{cli: c.cli, bucket: b}, nil
}

type cosFS struct {
	cli    *cos.Client
	bucket *cos.BucketGetResult
}

func (c *cosFS) Open(name string) (fs.File, error) {
	// TODO implement me
	panic("implement me")
}

func (c *cosFS) Stat(name string) (fs.FileInfo, error) {
	// TODO implement me
	panic("implement me")
}

func (c *cosFS) ReadDir(name string) ([]fs.DirEntry, error) {
	entries := make([]fs.DirEntry, 0, len(c.bucket.Contents))
	for _, v := range c.bucket.Contents {
		entries = append(entries, &objectFile{v})
	}
	return entries, nil
}

func (c *cosFS) ReadFile(name string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

type objectFile struct {
	obj cos.Object
}

func (o *objectFile) Type() fs.FileMode          { return o.Mode() }
func (o *objectFile) Info() (fs.FileInfo, error) { return o, nil }
func (o *objectFile) Name() string               { return o.obj.Key }
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
