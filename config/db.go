package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/anuraaaa/threadgo/domain"
)

func MustOpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("threadgo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening db: %v", err)
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.Comment{}, &domain.Like{}); err != nil {
		log.Fatalf("auto-migrate: %v", err)
	}

	return db
}
