package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	RoleID   uint   `gorm:"not null;DEFAULT:2" json:"role_id"`
	Username string `gorm:"size:32;not null;unique" json:"username"`
	Email    string `gorm:"size:80;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	FullName string `gorm:"size:100;" json:"fullname"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"role"`
}
