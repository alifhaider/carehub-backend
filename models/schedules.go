package models

import (
	"time"

	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	DoctorID      uint      `json:"doctor_id"`
	Doctor        Doctor    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	StartTime     time.Time `json:"start_time" gorm:"default:CURRENT_TIMESTAMP"`
	EndTime       time.Time `json:"end_time" gorm:"default:CURRENT_TIMESTAMP"`
	Location      Location  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"location"`
	LocationID    uint64    `gorm:"not null" json:"location_id"`
	Bookings      []Booking `gorm:"foreignKey:ScheduleID"`
	SerialFee     int       `json:"serial_fee"`
	VisitFee      int       `json:"visit_fee"`
	Discount      int       `json:"discount"`
	DepositAmount int       `json:"deposit_amount"`
	// MaxAppointment represents the maximum number of appointments allowed.
	// It is serialized to JSON as "max_appointment" and has a default value of 10 in the database.
	MaxAppointment int `json:"max_appointment" gorm:"default:10"`
}

// Save schedule details
func (schedule *Schedule) Save() (*Schedule, error) {
	err := database.Db.Create(&schedule).Error
	if err != nil {
		return &Schedule{}, err
	}
	return schedule, nil
}

// get schedules by doctor id
func GetSchedulesByDoctorId(doctorId uint, Schedule *[]Schedule) (err error) {
	err = database.Db.Where("doctor_id=?", doctorId).Find(Schedule).Error
	if err != nil {
		return err
	}
	return nil
}

// get schedules by location id
func GetSchedulesByLocationId(locationId uint, Schedule *[]Schedule) (err error) {
	err = database.Db.Where("location_id=?", locationId).Find(Schedule).Error
	if err != nil {
		return err
	}
	return nil
}

// get schedules by doctor id and location id
func GetSchedulesByDoctorIdAndLocationId(doctorId, locationId uint, Schedule *[]Schedule) (err error) {
	err = database.Db.Where("doctor_id=? AND location_id=?", doctorId, locationId).Find(Schedule).Error
	if err != nil {
		return err
	}
	return nil
}

// get schedule by id
func GetScheduleById(id uint) (Schedule, error) {
	var schedule Schedule
	err := database.Db.Where("id=?", id).Find(&schedule).Error
	if err != nil {
		return Schedule{}, err
	}
	return schedule, nil
}

// update schedule details
func (schedule *Schedule) Update() (*Schedule, error) {
	err := database.Db.Save(&schedule).Error
	if err != nil {
		return &Schedule{}, err
	}
	return schedule, nil
}

// delete schedule
func (schedule *Schedule) Delete() (err error) {
	err = database.Db.Delete(&schedule).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all schedules
func GetSchedules(Schedule *[]Schedule) (err error) {
	err = database.Db.Find(Schedule).Error
	if err != nil {
		return err
	}
	return nil
}
