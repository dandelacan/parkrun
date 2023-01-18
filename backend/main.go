package main

import (
	"fmt"
	"net/http"

	"parkrun/backend/controllers"
	"parkrun/backend/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	host     = "parkrundb.c5bt6dubggbv.eu-west-2.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "sorry, we couldn't find what you were looking for")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("main")

	// us, err := models.NewUserService(psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer us.Close()

	trs, err := models.NewTimesService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer trs.Close()
	fmt.Println("conected to db")
	// us.AutoMigrate()

	// usersController := controllers.NewUsers(us)
	timeRecordController := controllers.NewTimesService(trs)

	mux := mux.NewRouter()

	// mux.HandleFunc("/users", usersController.ViewUsers)
	mux.HandleFunc("/times", timeRecordController.ViewTimes).Methods("GET")
	mux.HandleFunc("/times", timeRecordController.AddTimes).Methods("POST")
	mux.HandleFunc("/", timeRecordController.AddTimes)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":3000", handler)
}
