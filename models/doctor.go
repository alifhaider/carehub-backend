package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	PhoneNumber    *string     `gorm:"size:255;unique" json:"phone_number"`
	ProfilePicture *string     `gorm:"size:255;" json:"profile_picture"`
	Specialties    []Specialty `gorm:"many2many:doctor_specialties;"`
	UserID         uint        `gorm:"not null" json:"user_id"`
	User           User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Education      []Education `gorm:"foreignKey:DoctorID"`
	Bio            string      `gorm:"size:255;" json:"bio"`
	Schedules      []Schedule  `gorm:"foreignKey:DoctorID"`
	Bookings       []Booking   `gorm:"foreignKey:DoctorID"`
	Rating         int         `gorm:"not null" json:"rating"`
	Reviews        []Review    `gorm:"foreignKey:DoctorID"`
}

// Save doctor details
func (doctor *Doctor) Save() (*Doctor, error) {
	err := database.Db.Create(&doctor).Error
	if err != nil {
		return &Doctor{}, err
	}
	return doctor, nil
}

// Get all doctors
func GetDoctors(Doctor *[]Doctor) (err error) {
	err = database.Db.Find(Doctor).Error
	if err != nil {
		return err
	}
	return nil
}

// Get doctor by id
func GetDoctorById(id uint) (Doctor, error) {
	var doctor Doctor
	err := database.Db.Where("id=?", id).Find(&doctor).Error
	if err != nil {
		return Doctor{}, err
	}
	return doctor, nil
}

// Get doctor by user id
func GetDoctorByUserId(userId uint) (Doctor, error) {
	var doctor Doctor
	err := database.Db.Where("user_id=?", userId).Find(&doctor).Error
	if err != nil {
		return Doctor{}, err
	}
	return doctor, nil
}

// update doctor details
func (doctor *Doctor) Update() (*Doctor, error) {
	err := database.Db.Save(&doctor).Error
	if err != nil {
		return &Doctor{}, err
	}
	return doctor, nil
}
