package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Conn *gorm.DB
}

func NewDbConnection() (*DbConnection, error) {
	//dsn := "host=localhost user=vets_backend password=vets dbname=vets port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &DbConnection{
		db,
	}, err

}
