package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating   int     `gorm:"default:0" json:"rating"`
	Comment  *string `gorm:"size:255" json:"comment"`
	DoctorID uint64  `json:"doctor_id"`
	Doctor   Doctor  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	UserID   uint64  `json:"user_id" gorm:"index"`
	User     User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// Save review details
func (review *Review) Save() (*Review, error) {
	err := database.Db.Create(&review).Error
	if err != nil {
		return &Review{}, err
	}
	return review, nil
}

// Get all reviews for a doctor
func GetReviewsByDoctorId(doctorId uint, Review *[]Review) (err error) {
	err = database.Db.Where("doctor_id=?", doctorId).Find(Review).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all reviews for a user
func GetReviewsByUserId(userId uint, Review *[]Review) (err error) {
	err = database.Db.Where("user_id=?", userId).Find(Review).Error
	if err != nil {
		return err
	}
	return nil
}
