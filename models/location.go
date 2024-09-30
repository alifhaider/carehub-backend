package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	City      string     `json:"city"`
	State     string     `json:"state"`
	ZipCode   string     `json:"zipCode"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Schedules []Schedule `gorm:"foreignKey:LocationID" json:"schedules,omitempty"`
	Bookings  []Booking  `gorm:"foreignKey:LocationID" json:"bookings,omitempty"`
}
