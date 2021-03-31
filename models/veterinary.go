package models

import "gorm.io/gorm"

type Veterinary struct {
	gorm.Model
	Title       string
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Phone       string `gorm:"unique"`
	Address     string
	Latitude    float64
	Longitude   float64
	VetServices []VetService
}
