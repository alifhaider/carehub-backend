package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	DoctorID      uint     `gorm:"not null" json:"doctor_id"`
	Doctor        Doctor   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	UserID        uint     `gorm:"not null" json:"user_id"`
	User          User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	ScheduleID    uint     `gorm:"not null" json:"schedule_id"`
	Schedule      Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"schedule"`
	LocationID    uint     `gorm:"not null" json:"location_id"`
	Location      Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"location"`
	BookingStatus string   `gorm:"size:255;" json:"booking_status"`
}

// Save booking details
func (booking *Booking) Save() (*Booking, error) {
	err := database.Db.Create(&booking).Error
	if err != nil {
		return &Booking{}, err
	}
	return booking, nil
}

// Get all bookings
func GetBookings(Booking *[]Booking) (err error) {
	err = database.Db.Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get booking by id
func GetBookingById(id uint) (Booking, error) {
	var booking Booking
	err := database.Db.Where("id=?", id).Find(&booking).Error
	if err != nil {
		return Booking{}, err
	}
	return booking, nil
}

// update booking details
func (booking *Booking) Update() (*Booking, error) {
	err := database.Db.Save(&booking).Error
	if err != nil {
		return &Booking{}, err
	}
	return booking, nil
}

// delete booking
func (booking *Booking) Delete() (err error) {
	err = database.Db.Delete(&booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all bookings by user id
func GetBookingsByUserId(userId uint, Booking *[]Booking) (err error) {
	err = database.Db.Where("user_id=?", userId).Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all bookings by doctor id
func GetBookingsByDoctorId(doctorId uint, Booking *[]Booking) (err error) {
	err = database.Db.Where("doctor_id=?", doctorId).Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all bookings by schedule id
func GetBookingsByScheduleId(scheduleId uint, Booking *[]Booking) (err error) {
	err = database.Db.Where("schedule_id=?", scheduleId).Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all bookings by location id
func GetBookingsByLocationId(locationId uint, Booking *[]Booking) (err error) {
	err = database.Db.Where("location_id=?", locationId).Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all bookings by booking status
func GetBookingsByBookingStatus(bookingStatus string, Booking *[]Booking) (err error) {
	err = database.Db.Where("booking_status=?", bookingStatus).Find(Booking).Error
	if err != nil {
		return err
	}
	return nil
}
