package models

import "gorm.io/gorm"

type Location struct {
	longitude float64
	latitude  float64
}

type Veterinary struct {
	gorm.Model
	Title     string
	FirstName string
	LastName  string
	Summary   string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Address   string
	// Latitude  float64
	// Longitude float64
	Location    Location `gorm:"type:point"`
	VetServices []VetService
	Sessions    []Session
}
