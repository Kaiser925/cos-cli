package oss

import (
	"io/fs"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// COSObject wrappers cos.Object.
// It implements the fs.FileInfo.
type COSObject struct {
	obj *cos.Object
}

// COSObjects is the slice of *COSObject
type COSObjects []*COSObject

func NewCOSObjects(objs []cos.Object) COSObjects {
	cobjs := make([]*COSObject, 0, len(objs))
	for _, obj := range objs {
		cobjs = append(cobjs, NewCOSObject(&obj))
	}
	return cobjs
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
		return fs.ModeDir | 0o444
	}
	return 0o444
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
