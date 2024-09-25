package models

import "gorm.io/gorm"

type Doctor struct {
    gorm.Model
    Name       string
    Specialty  string
    Location   string
    Users      []User `gorm:"many2many:user_doctors;"`
}