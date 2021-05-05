package models

import "gorm.io/gorm"

type Farmer struct {
	gorm.Model `faker:"-"`
	FirstName  string    `faker:"first_name"`
	LastName   string    `faker:"last_name"`
	Email      string    `gorm:"unique" faker:"email"`
	Phone      string    `gorm:"unique" faker:"phone_number"`
	Address    string    `faker:"-"`
	Latitude   float64   `faker:"lat"`
	Longitude  float64   `faker:"long"`
	Sessions   []Session `faker:"-"`
}
