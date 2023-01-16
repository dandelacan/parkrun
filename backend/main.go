package main

import (
	"fmt"
	"net/http"

	"parkrun/backend/controllers"
	"parkrun/backend/models"

	"github.com/gorilla/handlers"
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

	trs, err := models.NewTimesService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer trs.Close()
	fmt.Println("conected to db")
	us.AutoMigrate()

	usersController := controllers.NewUsers(us)
	timeRecordController := controllers.NewTimesService(trs)

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Origin", "Content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	func corsMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleware", r.Method)
	   
	   
		if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		return
		}
	   
	   
		next.ServeHTTP(w, r)
		log.Println("Executing middleware again")
		})
	   }

	r := mux.NewRouter()
	r.Use(cors)
	r.Use()
	r.HandleFunc("/users", usersController.ViewUsers)
	r.HandleFunc("/times", timeRecordController.ViewTimes).Methods("GET")
	r.HandleFunc("/times", timeRecordController.AddTimes).Methods("POST")

	http.ListenAndServe(":3000", corsMiddleware(r))
}
