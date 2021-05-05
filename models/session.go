package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model         `faker:"-"`
	Date               time.Time            `faker:"timet"`
	Latitude           float64              `faker:"lat"`
	Longitude          float64              `faker:"long"`
	FarmerID           uint                 `faker:"boundary_start=2, boundary_end=19"`
	VeterinaryID       uint                 `faker:"boundary_start=2, boundary_end=24"`
	VetServiceSessions []*VetServiceSession `faker:"-"`
}
