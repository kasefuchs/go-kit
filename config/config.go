package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	delimiter  = "."
	underscore = "_"
)

// Config wrapper for generic type T
type Config[T any] struct {
	k *koanf.Koanf
}

// New creates config instance
func New[T any]() *Config[T] {
	return &Config[T]{
		k: koanf.New(delimiter),
	}
}

// LoadMap loads config from map
func (c *Config[T]) LoadMap(mp map[string]any) error {
	if err := c.k.Load(confmap.Provider(mp, delimiter), nil); err != nil {
		return fmt.Errorf("failed to load config values")
	}

	return nil
}

// LoadEnv transforms env vars and loads config
func (c *Config[T]) LoadEnv(prefix string) error {
	prefix = strings.ToUpper(prefix) + underscore
	provider := env.Provider(prefix, delimiter, func(s string) string {
		s = strings.TrimPrefix(s, prefix)
		s = strings.ToLower(s)
		return strings.ReplaceAll(s, underscore, delimiter)
	})

	if err := c.k.Load(provider, nil); err != nil {
		return fmt.Errorf("failed to load env: %w", err)
	}

	return nil
}

// LoadFile loads config from file
func (c *Config[T]) LoadFile(path string, parser koanf.Parser) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of config file: %w", err)
	}

	provider := file.Provider(absPath)
	if err := c.k.Load(provider, parser); err != nil {
		return fmt.Errorf("failed to load config file at %s: %w", absPath, err)
	}

	return nil
}

// Build unmarshals data into struct T
func (c *Config[T]) Build() (T, error) {
	var v T
	if err := c.k.Unmarshal("", &v); err != nil {
		return v, err
	}

	return v, nil
}

// String returns debug string
func (c *Config[T]) String() string {
	return c.k.Sprint()
}
