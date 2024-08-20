package db

import (
	"fmt"
	"hehe/entity"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

// create new connection
func NewDB() (*DB, error) {
	// buat koneksi ke sqlite
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return nil, err
	}

	return &DB{Db: db}, nil
}

// get db
func GetDB() (*DB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// migration
func (d *DB) Migrate() error {
	// migrate table
	if err := d.Db.AutoMigrate(&entity.User{}, &entity.CodeSubscription{}, &entity.Berapakali{}); err != nil {
		return err
	}
	return nil
}
