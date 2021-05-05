package models

import (
	"gorm.io/gorm"
)

type VetServiceSession struct {
	gorm.Model   `faker:"-"`
	VetServiceID uint `faker:"boundary_start=29, boundary_end=57"`
	SessionID    uint `faker:"boundary_start=25, boundary_end=36"`
	Units        uint `faker:"boundary_start=1, boundary_end=500"`
}
