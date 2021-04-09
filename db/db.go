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
	//dsn := "postgres://rhdktlxtufaaoe:22971c5d488da6a8e27011251b24f2d99eff8af796deeaaec49b13efb557330c@ec2-52-1-115-6.compute-1.amazonaws.com:5432/da5kla8im8n293"
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &DbConnection{
		db,
	}, err

}
