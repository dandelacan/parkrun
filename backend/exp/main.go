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

	us, err := models.NewTimesService(psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("ecp connected to db")
	defer us.Close()

	dan := models.TimeRecord{
		Time:      30000,
		BarcodeID: 33,
	}
	val := models.TimeRecord{
		Time:      20000,
		BarcodeID: 43,
	}
	bob := models.TimeRecord{
		Time:      25000,
		BarcodeID: 25,
	}
	us.DestructiveReset()
	us.Create(&dan)
	us.Create(&val)
	us.Create(&bob)
}
