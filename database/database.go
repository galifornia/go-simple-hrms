package database

import (
	"github.com/galifornia/go-simple-hrms/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("hrms.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&types.Employee{})

	return DB
}
