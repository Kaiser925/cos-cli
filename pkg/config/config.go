package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	SecretKey  string `json:"secret_key"`
}

// New returns the default *Config.
func New() *Config {
	return &Config{
		Version: ClientVersion,
		Aliases: make(map[string]*AliasConfig),
	}
}

// WriteTo writes config data to w.
func (c *Config) WriteTo(w io.Writer) (int64, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	b, err := json.Marshal(c)
	if err != nil {
		return 0, fmt.Errorf("couldn't marshal config: %w", err)
	}
	return bytes.NewBuffer(b).WriteTo(w)
}

// SetAlias sets alias config
func (c *Config) SetAlias(alias string, cfg *AliasConfig) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.Aliases[alias] = cfg
}

// Save saves config to file.
func Save(c *Config, name string) error {
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

// Load loads config from file.
func Load(name string) (*Config, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cfg *Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
