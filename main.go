package main

import (
	"backend_development_go/config"
	"backend_development_go/handlers"
	"backend_development_go/middleware"
	"backend_development_go/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Book{}, &models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	api.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")

	admin := api.NewRoute().Subrouter()
	admin.Use(middleware.AdminOnlyMiddleware)

	admin.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	admin.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	admin.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
