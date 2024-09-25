package models

import "gorm.io/gorm"

type Doctor struct {
    gorm.Model
    FullName       *string  `gorm:"type:varchar(255)" json:"fullName,omitempty"`
		PhoneNumber    *string  `gorm:"type:varchar(255)" json:"phoneNumber,omitempty"`
		ProfilePicture  *string  `gorm:"type:varchar(255)" json:"profilePicture,omitempty"`
		Specialties     []Specialty  `gorm:"foreignKey:DoctorID" json:"specialties,omitempty"`
		UserID          uint             `gorm:"unique" json:"userId"`
		User            User               `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
    Education       []Education        `gorm:"foreignKey:DoctorID" json:"education,omitempty"`
    Bio             string             `json:"bio"`
    Schedules       []Schedule         `gorm:"foreignKey:DoctorID" json:"schedules,omitempty"`
    Bookings        []Booking          `gorm:"foreignKey:DoctorID" json:"bookings,omitempty"`
    Rating          int                `gorm:"default:0" json:"rating"`
    Reviews         []Review           `gorm:"foreignKey:DoctorID" json:"reviews,omitempty"`
    Users      []User `gorm:"many2many:user_doctors;"`
}