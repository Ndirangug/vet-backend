package models

import "gorm.io/gorm"

type Veterinary struct {
	gorm.Model
	Title       string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Address     string
	Latitude    float64
	Longitude   float64
	VetServices []VetService
}
