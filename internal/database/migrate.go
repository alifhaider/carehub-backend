package database

import (
	"log"

	"github.com/alifhaider/carehub-backend/internal/models"
)

func MigrateDB() {
    err := DB.AutoMigrate(
			&models.User{},
        &models.Doctor{},
				&models.Password{},
				&models.Education{},
				&models.Specialty{},
				&models.Schedule{},
				&models.Booking{},
				&models.Review{},
				&models.Location{},
				
    )
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
}