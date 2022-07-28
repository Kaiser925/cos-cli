package config

import (
	"io/ioutil"
	"os"
	"testing"
)

var tmp = "tmp.json"

func TestSave(t *testing.T) {
	defer os.Remove(tmp)
	testcases := []struct {
		config  *Config
		content string
	}{
		{
			config: &Config{
				Version: ClientVersion,
				Aliases: map[string]*AliasConfig{},
			},
			content: `{
  "version": "0.1",
  "aliases": {}
}`,
		},
		{
			config: &Config{
				Version: ClientVersion,
				Aliases: map[string]*AliasConfig{
					"a": {
						BucketName: "bucket",
						Region:     "region",
						SecretID:   "id",
						SecretKey:  "key",
					},
				},
			},
			content: `{
  "version": "0.1",
  "aliases": {
    "a": {
      "bucketName": "bucket",
      "region": "region",
      "secretID": "id",
      "secretKey": "key"
    }
  }
}`,
		},
	}

	for i, test := range testcases {
		err := Save(test.config, tmp)
		if err != nil {
			t.Fatalf("test case %d failed, got error %v", i+1, err)
		}
		b, err := ioutil.ReadFile(tmp)
		if err != nil {
			t.Fatalf("test case %d failed, got error %v", i+1, err)
		}
		if got := string(b); got != test.content {
			t.Fatalf("test case %d failed, want \n%s \ngot\n %s", i+1, test.content, got)
		}
	}
}

func TestLoadOrInit(t *testing.T) {
	defer os.Remove(tmp)
	cfg, err := LoadOrInit(tmp)
	if err != nil {
		t.Fatalf(err.Error())
	}
	want := New()
	if !equal(cfg, want) {
		t.Fatalf("want %+v got %+v", want, cfg)
	}

	content := `{
  "version": "0.1",
  "aliases": {}
}`

	b, err := ioutil.ReadFile(tmp)
	if err != nil {
		t.Fatalf("read tmp file failed, got error %v", err)
	}

	if got := string(b); got != content {
		t.Fatalf("file content not right, want %s got %s", content, got)
	}
}

func equal(x *Config, y *Config) bool {
	if x.Version != y.Version {
		return false
	}
	for k, v := range x.Aliases {
		if y.Aliases[k] != v {
			return false
		}
	}
	return true
}
