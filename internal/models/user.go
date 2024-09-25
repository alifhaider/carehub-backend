package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    UserName string    `gorm:"unique;type:varchar(255)" json:"username"`
    Email    string    `gorm:"unique;type:varchar(255)" json:"email"`
    Password *Password `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"password,omitempty"`
}

type Password struct {
    gorm.Model
    Hash   string `gorm:"type:text" json:"hash"`
    UserID uint   `gorm:"uniqueIndex" json:"userId"`
    User   User   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"user"`
}
