package models

import (
	"gorm.io/gorm"
)

type VetServiceSession struct {
	gorm.Model
	VetServiceID uint
	SessionID    uint
	Units        uint
}
