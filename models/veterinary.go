package models

import "gorm.io/gorm"

type Location struct {
	Longitude float64
	Latitude  float64
}

type Veterinary struct {
	gorm.Model
	Title       string
	FirstName   string
	LastName    string
	Summary     string
	Email       string `gorm:"unique"`
	Phone       string `gorm:"unique"`
	Address     string
	Latitude    float64
	Longitude   float64
	VetServices []VetService
	Sessions    []Session
}
