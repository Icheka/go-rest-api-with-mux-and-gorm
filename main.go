package main

import (
	"go-rest-api-with-mux-and-gorm/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initMultiplexer() {
	router := mux.NewRouter()

	router.HandleFunc("/users", database.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", database.GetUser).Methods("GET")
	router.HandleFunc("/users", database.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", database.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", database.DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	database.InitialMigration()
	initMultiplexer()
}
