package config

import (
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// Connect establishes a connection to the database
func Connect() {
	// Update the connection string
	d, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err) // Exit the program if there's a connection error
	}
	db = d
}

// GETDB returns the database connection instance
func GETDB() *gorm.DB {
	return db
}
