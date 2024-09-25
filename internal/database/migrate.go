package database

import (
	"log"

	"github.com/alifhaider/carehub-backend/internal/models"
)

func MigrateDB() {
    err := DB.AutoMigrate(
        &models.Doctor{}, // Include other models here as necessary
    )
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
}