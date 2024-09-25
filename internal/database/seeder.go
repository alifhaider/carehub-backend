package database

import (
	"log"

	"github.com/alifhaider/carehub-backend/internal/models"
)

func SeedDB() {
    doctors := []models.Doctor{
        {Name: "Dr. Smith", Specialty: "Cardiology", Location: "New York"},
        {Name: "Dr. Alice", Specialty: "Pediatrics", Location: "Los Angeles"},
    }
		println("Seeding doctors...")
    for _, doctor := range doctors {
        if err := DB.Create(&doctor).Error; err != nil {
            log.Println("Failed to seed doctor:", doctor.Name)
        }
    }
}