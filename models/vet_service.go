package models

import "gorm.io/gorm"

type VetService struct {
	gorm.Model
	VeterinaryID       uint
	ServiceID          uint
	Unit               string
	ChargePerUnit      float64
	VetServiceSessions []VetServiceSession
}
