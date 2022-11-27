package config

import (
	"os"
	"path"
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
						URL:       "https://bucket.cos.region.myqcloud.com",
						SecretID:  "id",
						SecretKey: "key",
					},
				},
			},
			content: `{
  "version": "0.1",
  "aliases": {
    "a": {
      "url": "https://bucket.cos.region.myqcloud.com",
      "secretID": "id",
      "secretKey": "key"
    }
  }
}`,
		},
	}

	for i, test := range testcases {
		err := test.config.Save(tmp)
		if err != nil {
			t.Fatalf("test case %d failed, got error %v", i+1, err)
		}
		b, err := os.ReadFile(tmp)
		if err != nil {
			t.Fatalf("test case %d failed, got error %v", i+1, err)
		}
		if got := string(b); got != test.content {
			t.Fatalf("test case %d failed, want \n%s \ngot\n %s", i+1, test.content, got)
		}
	}
}

func TestLoadOrInit(t *testing.T) {
	tmp := path.Join(os.TempDir(), "config.json")
	defer os.Remove(tmp)
	err := LoadOrInit(tmp)
	if err != nil {
		t.Fatalf(err.Error())
	}
	want := New()
	if !equal(config, want) {
		t.Fatalf("want %+v got %+v", want, config)
	}

	content := `{
  "version": "0.1",
  "aliases": {}
}`

	b, err := os.ReadFile(tmp)
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

func TestConfig_RemoveAlias(t *testing.T) {
	SetAlias("alias", nil)
	_, ok := GetAlias("alias")
	if !ok {
		t.Error("got alias failed, want ture, got false")
	}
	RemoveAlias("alias")
	_, ok = Default().GetAlias("alias")
	if ok {
		t.Error("remove alias failed")
	}
	// remove alias not exist
	RemoveAlias("alias")
}
