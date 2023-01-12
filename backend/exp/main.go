package main

import (
	"fmt"
	"parkrun/backend/models"
)

const (
	host     = "parkrundb.c5bt6dubggbv.eu-west-2.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("ecp connected to db")
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
