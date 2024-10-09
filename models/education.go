package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	DoctorID    uint   `gorm:"not null" json:"doctor_id"`
	Doctor      Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	Institution string `gorm:"size:255;" json:"institution"`
	Degree      string `gorm:"size:255;" json:"degree"`
	PassedYear  uint   `gorm:"not null" json:"passed_year"`
}

// Save education details
func (education *Education) Save() (*Education, error) {
	err := database.Db.Create(&education).Error
	if err != nil {
		return &Education{}, err
	}
	return education, nil
}

// update education details
func (education *Education) Update() (*Education, error) {
	err := database.Db.Save(&education).Error
	if err != nil {
		return &Education{}, err
	}
	return education, nil
}
