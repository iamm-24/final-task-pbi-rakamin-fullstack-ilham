package database

import (
	"rakamin/app"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(&app.User{}, &app.Photo{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}