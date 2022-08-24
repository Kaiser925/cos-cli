package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
)

const (
	ClientVersion = "0.1"
)

// Config is cos-cli config
type Config struct {
	Version string                  `json:"version"`
	Aliases map[string]*AliasConfig `json:"aliases"`
	mux     sync.Mutex
}

// AliasConfig is a COS bucket config.
type AliasConfig struct {
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
}

// New returns the default *Config.
func New() *Config {
	return &Config{
		Version: ClientVersion,
		Aliases: make(map[string]*AliasConfig),
	}
}

// config is inner default config
var config = New()

// Default returns the default config.
func Default() *Config { return config }

// WriteTo writes config data to w.
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return 0, fmt.Errorf("couldn't marshal config: %w", err)
	}
	return bytes.NewBuffer(b).WriteTo(w)
}

// SetAlias sets alias config, if alias exists, it will be updated
// and return true.
func (c *Config) SetAlias(alias string, cfg *AliasConfig) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.Aliases[alias]
	c.Aliases[alias] = cfg
	return ok
}

// Save saves config to file.
func (c *Config) Save(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = c.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Load(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}

func (c *Config) LoadOrInit(name string) error {
	// create base dir if it not exists.
	dir := path.Dir(name)
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.MkdirAll(dir, 0750); err != nil {
			return err
		}
	}

	// stat config file, if not exists, create config file.
	if _, err := os.Stat(name); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := c.Save(name); err != nil {
			return err
		}
	}

	return c.Load(name)
}

// Save saves default config to file.
func Save(name string) error {
	return config.Save(name)
}

// Load loads default config from file.
func Load(name string) (*Config, error) {
	if err := config.Load(name); err != nil {
		return nil, err
	}
	return config, nil
}

// SetAlias sets an alias for default Config.
func SetAlias(alias string, cfg *AliasConfig) bool {
	return config.SetAlias(alias, cfg)
}

// LoadOrInit loads default config from given file.
// If file not exists, it will create file, and returns default config.
func LoadOrInit(name string) error {
	return config.LoadOrInit(name)
}
