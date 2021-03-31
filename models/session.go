package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	date               time.Time
	ServiceID          uint
	Unit               string
	ChargePerUnit      float64
	Latitude           float64
	Longitude          float64
	FarmerID           uint
	VetServiceSessions []VetServiceSession
}
