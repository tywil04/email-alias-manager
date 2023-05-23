package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE_PATH")), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&User{}, &Alias{})

	DB = db
	return nil
}
