package models

import "gorm.io/gorm"

type Location struct {
	Longitude float64
	Latitude  float64
}

type Veterinary struct {
	gorm.Model  `faker:"-"`
	Title       string       `faker:"oneof: Dr, Mr, Madam"`
	FirstName   string       `faker:"first_name"`
	LastName    string       `faker:"first_name"`
	Summary     string       `faker:"oneof: livestock veterinarian, livestock physician, pet veterinarian"`
	Email       string       `gorm:"unique" faker:"email"`
	Phone       string       `gorm:"unique" faker:"phone_number"`
	Address     string       `faker:"-"`
	Latitude    float64      `faker:"lat"`
	Longitude   float64      `faker:"long"`
	VetServices []VetService `faker:"-"`
	Sessions    []Session    `faker:"-"`
}
