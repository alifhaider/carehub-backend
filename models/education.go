package models

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	DoctorID    uint   `json:"doctorId"`
	Doctor      Doctor `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`
	Institution string `gorm:"type:varchar(255)" json:"institution"`
	Degree      string `gorm:"type:varchar(255)" json:"degree"`
	PassedYear  uint   `json:"passedYear"`
}
