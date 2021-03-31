package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Conn *gorm.DB
}

func NewDbConnection() (*DbConnection, error) {
	dsn := "host=localhost user=vets_backend password=vets dbname=vets port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &DbConnection{
		db,
	}, err

}
