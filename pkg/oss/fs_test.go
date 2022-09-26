package oss

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/fs"
	"testing"
	"time"
)

func TestCOSObject(t *testing.T) {
	now := time.Now()
	nt, _ := time.Parse(time.RFC3339, now.Format(time.RFC3339))
	tests := []struct {
		object  cos.Object
		name    string
		isDir   bool
		size    int64
		mode    fs.FileMode
		modTime time.Time
	}{
		{
			object:  cos.Object{Key: "key", Size: 654, LastModified: now.Format(time.RFC3339)},
			name:    "key",
			size:    654,
			mode:    0444,
			modTime: nt,
			isDir:   false,
		},
		{
			object:  cos.Object{Key: "bucket/is-dir", Size: 0, LastModified: now.Format(time.RFC3339)},
			name:    "bucket/is-dir",
			size:    0,
			mode:    0444 | fs.ModeDir,
			modTime: nt,
			isDir:   true,
		},
	}

	for i, test := range tests {
		obj := NewCOSObject(&test.object)
		if obj.Name() != test.name {
			t.Errorf("test case %d failed(Name), want %s got %s", i+1, test.name, obj.Name())
		}
		if obj.IsDir() != test.isDir {
			t.Errorf("test case %d failed(IsDir), want %v got  %v", i+1, test.isDir, obj.IsDir())
		}
		if obj.Size() != test.size {
			t.Errorf("test case %d failed(Size), want %v got  %v", i+1, test.size, obj.Size())
		}
		if obj.ModTime() != test.modTime {
			t.Errorf("test case %d failed(ModTime), want %v got  %v", i+1, test.modTime, obj.ModTime())
		}
		if obj.Mode() != test.mode {
			t.Errorf("test case %d failed(Mod), want %v got  %v", i+1, test.mode, obj.Mode())
		}
	}
}
