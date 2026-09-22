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
	koanfDelimiter = "."
	envDelimiter   = "__"
)

var k = koanf.New(koanfDelimiter)

// Koanf returns global koanf instance
func Koanf() *koanf.Koanf {
	return k
}

// String returns debug string
func String() string {
	return k.Sprint()
}

// LoadMap loads config from map
func LoadMap(mp map[string]any) error {
	if err := k.Load(confmap.Provider(mp, koanfDelimiter), nil); err != nil {
		return fmt.Errorf("failed to load config values")
	}

	return nil
}

// LoadEnv transforms env vars and loads config
func LoadEnv(prefix string) error {
	prefix = strings.ToUpper(prefix) + envDelimiter
	provider := env.Provider(prefix, koanfDelimiter, func(s string) string {
		s = strings.TrimPrefix(s, prefix)
		s = strings.ToLower(s)

		return strings.ReplaceAll(s, envDelimiter, koanfDelimiter)
	})

	if err := k.Load(provider, nil); err != nil {
		return fmt.Errorf("failed to load env: %w", err)
	}

	return nil
}

// LoadFile loads config from file
func LoadFile(path string, parser koanf.Parser) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of config file: %w", err)
	}

	provider := file.Provider(absPath)
	if err := k.Load(provider, parser); err != nil {
		return fmt.Errorf("failed to load config file at %s: %w", absPath, err)
	}

	return nil
}

// Build unmarshals data into struct T
func Build[T any]() (T, error) {
	var v T
	if err := k.Unmarshal("", &v); err != nil {
		return v, err
	}

	return v, nil
}
