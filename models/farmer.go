package models

import "gorm.io/gorm"

type Farmer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Address   string
	Latitude  float64
	Longitude float64
	Sessions  []Session
}
