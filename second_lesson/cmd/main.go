package main

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/usmonali4/gotour/pkg/db"
	"github.com/usmonali4/gotour/pkg/handlers"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)


	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}