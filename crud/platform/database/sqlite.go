package database

import (
	"os"

	"github.com/Sadham-Hussian/go-learnings/crud/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLConnection function to connect db and return the connection
func SQLConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLLITE_DATABASE")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.Product{})

	return db, nil
}
