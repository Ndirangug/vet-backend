package models

import "gorm.io/gorm"

type VetService struct {
	gorm.Model         `faker:"-"`
	VeterinaryID       uint                `faker:"boundary_start=2, boundary_end=24"`
	ServiceID          uint                `faker:"boundary_start=2, boundary_end=9"`
	Unit               string              `faker:"oneof: cow, cow"`
	ChargePerUnit      float64             `faker:"amount"`
	VetServiceSessions []VetServiceSession `faker:"-"`
}
