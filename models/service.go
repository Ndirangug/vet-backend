package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name        string
	Description string
	VetServices []VetService
}
