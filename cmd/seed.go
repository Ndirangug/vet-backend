package cmd

import (
	"fmt"
	"reflect"
	"time"

	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/ndirangug/vets-backend/db"
	"github.com/ndirangug/vets-backend/models"
)

func Seed() {
	registerCustomGenerators()

	db, err := db.NewDbConnection()

	if err != nil {
		panic(err)
	}

	//db.Conn.AutoMigrate(&models.Vehicle{}, &models.Driver{}, &models.Trip{})

	//	TODO: first clear all data from tables then begin seeding

	farmers := []*models.Farmer{}
	veterinries := []*models.Veterinary{}
	sessions := []*models.Session{}
	vetServices := []*models.VetService{}
	vetServiceSessions := []*models.VetServiceSession{}

	faker.SetRandomMapAndSliceSize(30)

	err = faker.FakeData(&farmers)
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&veterinries)
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&sessions)
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&vetServices)
	if err != nil {
		panic(err)
	}
	err = faker.FakeData(&vetServiceSessions)
	if err != nil {
		panic(err)
	}

	//db.Conn.Create(farmers)
	//db.Conn.Create(veterinries)
	//db.Conn.Create(vetServices)
	//db.Conn.Create(sessions)
	db.Conn.Create(vetServiceSessions)

}

func registerCustomGenerators() {

	_ = faker.AddProvider("timet", func(v reflect.Value) (interface{}, error) {
		min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
		delta := max - min

		sec := rand.Int63n(delta) + min
		return time.Unix(sec, 0), nil
	})

	_ = faker.AddProvider("avatar", func(v reflect.Value) (interface{}, error) {
		return fmt.Sprintf("https://i.pravatar.cc/150?img=%d", rand.Intn(70)), nil
	})

}
