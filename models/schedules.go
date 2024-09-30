package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	DoctorID       uint      `json:"doctorId"`
	Doctor         Doctor    `gorm:"constraint:OnDelete:CASCADE;" json:"doctor"`
	StartTime      string    `json:"startTime"`
	EndTime        string    `json:"endTime"`
	Date           time.Time `json:"date"`
	Location       Location  `gorm:"constraint:OnDelete:CASCADE;" json:"location"`
	LocationID     uint      `json:"locationId"`
	Bookings       []Booking `gorm:"foreignKey:ScheduleID" json:"bookings,omitempty"`
	SerialFee      int       `json:"serialFee"`
	VisitFee       int       `json:"visitFee"`
	Discount       int       `json:"discount"`
	DepositAmount  int       `json:"depositAmount"`
	MaxAppointment int       `json:"maxAppointment"`
}
