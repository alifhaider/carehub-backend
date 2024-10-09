package models

import (
	"github.com/alifhaider/carehub-backend/database"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name      string     `gorm:"size:255;unique" json:"name"`
	Address   string     `gorm:"size:255" json:"address"`
	City      string     `gorm:"size:255" json:"city"`
	State     *string    `gorm:"size:255;" json:"state"`
	ZipCode   *string    `gorm:"size:255;" json:"zip_code"`
	Latitude  *float64   `json:"latitude"`
	Longitude *float64   `json:"longitude"`
	Schedules []Schedule `gorm:"foreignKey:SchedulesID"`
	Bookings  []Booking  `gorm:"foreignKey:BookingsID"`
}

// Save location details
func (location *Location) Save() (*Location, error) {
	err := database.Db.Create(&location).Error
	if err != nil {
		return &Location{}, err
	}
	return location, nil
}

// Get all locations
func GetLocations(Location *[]Location) (err error) {
	err = database.Db.Find(Location).Error
	if err != nil {
		return err
	}
	return nil
}

// Get location by id
func GetLocationById(id uint) (Location, error) {
	var location Location
	err := database.Db.Where("id=?", id).Find(&location).Error
	if err != nil {
		return Location{}, err
	}
	return location, nil
}

// update location details
func (location *Location) Update() (*Location, error) {
	err := database.Db.Save(&location).Error
	if err != nil {
		return &Location{}, err
	}
	return location, nil
}

// delete location
func (location *Location) Delete() (err error) {
	err = database.Db.Delete(&location).Error
	if err != nil {
		return err
	}
	return nil
}

// search location by name or address or city or state or zip code
func SearchLocation(search string, location *[]Location) (err error) {
	err = database.Db.Where("name LIKE ? OR address LIKE ? OR city LIKE ? OR state LIKE ? OR zip_code LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(location).Error
	if err != nil {
		return err
	}
	return nil
}
