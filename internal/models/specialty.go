package models

import "gorm.io/gorm"

type Specialty struct {
    gorm.Model
    Name string `gorm:"type:varchar(255)" json:"name"`
		Doctor Doctor `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`
		DoctorID uint `json:"doctorId"`
}
