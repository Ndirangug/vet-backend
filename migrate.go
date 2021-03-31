package main

import (
	"github.com/ndirangug/vets-backend/db"
	"github.com/ndirangug/vets-backend/models"
)

func main() {
	db, err := db.NewDbConnection()

	if err != nil {
		panic(err)
	}

	db.Conn.AutoMigrate(&models.Farmer{}, &models.Veterinary{}, &models.Service{}, &models.VetService{}, &models.Session{}, &models.VetServiceSession{})
}
