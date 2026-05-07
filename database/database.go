package database

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// openDialector creates a default GORM dialector based on the driver name
func openDialector(driver, dsn string) (gorm.Dialector, error) {
	switch driver {
	case "sqlite":
		return sqlite.Open(dsn), nil
	case "postgres":
		return postgres.Open(dsn), nil
	default:
		return nil, fmt.Errorf("unsupported driver '%s'", driver)
	}
}

// DB returns the global gorm db instance
func DB() *gorm.DB {
	return db
}

// Init sets up the global database connection
func Init(cfg Config) error {
	dialector, err := openDialector(cfg.Driver, cfg.DSN)
	gormCfg := &gorm.Config{}

	if cfg.Debug {
		gormCfg.Logger = newGormLogger()
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(dialector, gormCfg)
	return err
}

// Close closes the underlying database connection
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		return sqlDB.Close()
	}

	return nil
}

// WithContext changes current instance db's context to ctx
func WithContext(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

// Model specifies the model you would like to run db operations with
func Model(value any) *gorm.DB {
	return db.Model(value)
}

// Table specifies the table you would like to run db operations with
func Table(name string, args ...any) *gorm.DB {
	return db.Table(name, args...)
}

// Where adds conditions
func Where(query any, args ...any) *gorm.DB {
	return db.Where(query, args...)
}

// Create inserts value into database
func Create(value any) *gorm.DB {
	return db.Create(value)
}

// Save updates value in database. If value doesn't have primary key, will insert it
func Save(value any) *gorm.DB {
	return db.Save(value)
}

// First finds first record that matches given conditions
func First(dest any, conds ...any) *gorm.DB {
	return db.First(dest, conds...)
}

// Find finds records that match given conditions
func Find(dest any, conds ...any) *gorm.DB {
	return db.Find(dest, conds...)
}

// Delete deletes value match given conditions
func Delete(value any, conds ...any) *gorm.DB {
	return db.Delete(value, conds...)
}

// AutoMigrate runs auto migration for given models
func AutoMigrate(dst ...any) error {
	return db.AutoMigrate(dst...)
}

// Transaction executes a function within a transaction
func Transaction(fc func(tx *gorm.DB) error) error {
	return db.Transaction(fc)
}
