package models

import (
	"path/filepath"
	"sync"

	"github.com/kabaliserv/tmpfiles/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

var database *DB

// GetDB is function to Get or make connection in Database
func GetDB() *DB {
	var once sync.Once
	once.Do(func() {
		dbPath := filepath.Join(config.GetStorePath(), "/data/database.db")

		db, _ := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

		database = &DB{db}
	})
	return database
}

// InitDB : make singleton connection to database
func InitDB() error {
	var err error
	var once sync.Once
	once.Do(func() {
		dbPath := filepath.Join(config.GetStorePath(), "/data/database.db")

		db, err1 := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err1 != nil {
			err = err1
		}
		database = &DB{db}
	})
	if err != nil {
		return err
	}
	return nil
}
