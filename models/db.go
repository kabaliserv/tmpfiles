package models

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewSqliteDB - sqlite database
func NewSqliteDB(dataPath string) (*DB, error) {

	dbPath := filepath.Join(dataPath, "database.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
		// panic("failed to connect database")
	}

	return &DB{db}, nil
}
