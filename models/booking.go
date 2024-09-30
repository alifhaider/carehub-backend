package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	DoctorID      uint     `json:"doctorId"`
	Doctor        Doctor   `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`
	UserID        uint     `json:"userId"`
	User          User     `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
	ScheduleID    uint     `json:"scheduleId"`
	Schedule      Schedule `gorm:"constraint:OnDelete:CASCADE;" json:"schedule"`
	LocationID    uint     `json:"locationId"`
	Location      Location `gorm:"constraint:OnDelete:CASCADE;" json:"location"`
	BookingStatus string   `gorm:"type:varchar(255);default:'PENDING'" json:"bookingStatus"`
}
