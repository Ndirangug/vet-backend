package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Date               time.Time
	Latitude           float64
	Longitude          float64
	FarmerID           uint
	VeterinaryID       uint
	VetServiceSessions []*VetServiceSession
}
