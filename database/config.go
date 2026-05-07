package database

// Config holds database settings
type Config struct {
	// Driver is the database dialect
	Driver string `koanf:"driver"`

	// DSN is the connection string
	DSN string `koanf:"dsn"`

	// Debug enables GORM logs
	Debug bool `koanf:"debug"`
}
