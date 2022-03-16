package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/VladMasarik/ef-notes/src/models"
)

var DB *gorm.DB

func SetupDatabase() error {
	var err error
	if DB, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{}); err != nil {
		return err
	}
	return DB.AutoMigrate(&models.Note{})
}
