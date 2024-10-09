package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Specialty struct {
	gorm.Model
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
	Doctor      Doctor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	DoctorID    uint64 `json:"doctor_id"`
}

// Save specialty details
func (specialty *Specialty) Save() (*Specialty, error) {
	err := database.Db.Create(&specialty).Error
	if err != nil {
		return &Specialty{}, err
	}
	return specialty, nil
}

// Get all specialties for a doctor
func GetSpecialtiesByDoctorId(doctorId uint, Specialty *[]Specialty) (err error) {
	err = database.Db.Where("doctor_id=?", doctorId).Find(Specialty).Error
	if err != nil {
		return err
	}
	return nil
}

// update specialty details
func (specialty *Specialty) Update() (*Specialty, error) {
	err := database.Db.Save(&specialty).Error
	if err != nil {
		return &Specialty{}, err
	}
	return specialty, nil
}

// delete specialty
func (specialty *Specialty) Delete() (err error) {
	err = database.Db.Delete(&specialty).Error
	return err
}

// Get all specialties
func GetSpecialties(Specialty *[]Specialty) (err error) {
	err = database.Db.Find(Specialty).Error
	if err != nil {
		return err
	}
	return nil
}

// Get specialty by id
func GetSpecialty(Specialty *Specialty, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Specialty).Error
	if err != nil {
		return err
	}
	return nil
}

// Get specialty by name
func GetSpecialtyByName(Specialty *Specialty, name string) (err error) {
	err = database.Db.Where("name = ?", name).First(Specialty).Error
	if err != nil {
		return err
	}
	return nil
}

// Get specialty by doctor id
func GetSpecialtyByDoctorId(Specialty *Specialty, doctorId uint) (err error) {
	err = database.Db.Where("doctor_id = ?", doctorId).First(Specialty).Error
	if err != nil {
		return err
	}
	return nil
}
