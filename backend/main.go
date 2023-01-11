package main

import (
	"fmt"
	"net/http"

	"parkrun/backend/controllers"
	"parkrun/backend/models"

	"github.com/gorilla/mux"
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

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	usersController := controllers.NewUsers(us)

	r := mux.NewRouter()

	r.HandleFunc("/users", usersController.ViewUsers)
	http.ListenAndServe(":3000", r)
}
