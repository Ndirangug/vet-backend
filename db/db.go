package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	conn *gorm.DB
}

func NewDbConnection() (*DbConnection, error) {
	dsn := "host=localhost user=vets_backend password=vets dbname=vets port=9920 sslmode=disable TimeZone=Africa/Nairobi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &DbConnection{
		db,
	}, err

}
