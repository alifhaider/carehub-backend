package database

import (
	"log"

	"github.com/alifhaider/carehub-backend/internal/models"
	"gorm.io/gorm"
)
func SeedDB(db *gorm.DB) {
    // Step 1: Delete existing records
    if err := db.Where("1 = 1").Delete(&models.Doctor{}).Error; err != nil {
        log.Fatalf("Failed to delete existing doctors: %v", err)
    }
    if err := db.Where("1 = 1").Delete(&models.User{}).Error; err != nil {
        log.Fatalf("Failed to delete existing users: %v", err)
    }

		log.Println("Deleted existing records")

    // Step 2: Create new users
    users := []models.User{
        {
            UserName: "alice123",
            Email:    "alice@example.com",
            Password: &models.Password{Hash: "hashed_password"}, // Replace with actual hash
        },
        {
            UserName: "bob456",
            Email:    "bob@example.com",
            Password: &models.Password{Hash: "hashed_password"}, // Replace with actual hash
        },
    }

    // Create users and check for errors
    for i := range users {
        if err := db.Create(&users[i]).Error; err != nil {
            log.Fatalf("Failed to seed user: %s, Error: %v", users[i].UserName, err)
        }
    }

    // Step 3: Create new doctors
    doctors := []models.Doctor{
        {
            FullName:       stringPtr("Dr. Alice"),
            PhoneNumber:    stringPtr("123-456-7890"),
            ProfilePicture: stringPtr("url_to_profile_picture"),
            UserID:        users[0].ID, // Ensure this ID is set correctly
            Bio:           "Pediatrician with over 10 years of experience.",
            Rating:        5,
        },
        {
            FullName:       stringPtr("Dr. Bob"),
            PhoneNumber:    stringPtr("098-765-4321"),
            ProfilePicture: stringPtr("url_to_profile_picture"),
            UserID:        users[1].ID, // Ensure this ID is set correctly
            Bio:           "General practitioner focusing on family health.",
            Rating:        4,
        },
    }

    // Create doctors and check for errors
    for _, doctor := range doctors {
        if err := db.Create(&doctor).Error; err != nil {
            log.Fatalf("Failed to seed doctor: %s, Error: %v", *doctor.FullName, err)
        }
    }

    log.Println("Successfully seeded users and doctors!")
}


func stringPtr(s string) *string {
    return &s
}