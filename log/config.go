package log

// Config holds logger settings
type Config struct {
	// Level is the log level
	Level string `koanf:"level"`

	// Pretty enables human-readable output
	Pretty bool `koanf:"pretty"`
}
