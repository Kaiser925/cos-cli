package oss

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/fs"
	"time"
)

// COSObject wrappers cos.Object.
// It implements the fs.FileInfo.
type COSObject struct {
	obj *cos.Object
}

func NewCOSObject(obj *cos.Object) *COSObject {
	return &COSObject{
		obj: obj,
	}
}

func (c *COSObject) Name() string {
	return c.obj.Key
}

func (c *COSObject) Size() int64 {
	return c.obj.Size
}

func (c *COSObject) Mode() fs.FileMode {
	// Assume there is no write permission on object
	if c.IsDir() {
		return fs.ModeDir | 0444
	}
	return 0444
}

func (c *COSObject) ModTime() time.Time {
	t, _ := time.Parse(time.RFC3339, c.obj.LastModified)
	return t
}

func (c *COSObject) IsDir() bool {
	return c.obj.Size == 0
}

func (c *COSObject) Sys() any {
	return nil
}
