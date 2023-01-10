package main

import (
	"fmt"
	"parkrun/backend/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "gheyghey"
	dbname   = "parkrun"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	dan := models.User{
		Name:      "dan",
		BarcodeID: 33,
	}
	val := models.User{
		Name:      "val",
		BarcodeID: 43,
	}
	bob := models.User{
		BarcodeID: 53,
		Name:      "bob",
	}
	us.DestructiveReset()
	us.Create(&dan)
	us.Create(&val)
	us.Create(&bob)
}
