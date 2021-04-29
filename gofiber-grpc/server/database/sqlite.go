package database

import (
	"os"

	//"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/server/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLConnection function to connect db and return the connection
func SQLConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(ProductItem{})

	return db, nil
}
