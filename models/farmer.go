package models

import "gorm.io/gorm"

type Farmer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	Latitude  float64
	Longitude float64
	Sessions  []Session
}
