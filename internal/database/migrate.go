package database

import (
	"log"

	"github.com/alifhaider/carehub-backend/internal/models"
)

func MigrateDB() {
    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database", err)
    }

    log.Println("Database migrated!")
}
