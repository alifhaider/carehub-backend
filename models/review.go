package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
	DoctorID uint   `json:"doctorId"`
	Doctor   Doctor `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`
	UserID   uint   `json:"userId"`
	User     User   `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
}
