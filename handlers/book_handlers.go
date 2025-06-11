package handlers

import (
	"backend_development_go/config"
	"backend_development_go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	config.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Get book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	config.DB.First(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

// Add a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

// update a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book []models.Book
	config.DB.First(&book, params["id"])
	json.NewDecoder(r.Body).Decode(&book)
	config.DB.Save(&book)
	json.NewEncoder(w).Encode(book)

}

// delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book []models.Book
	config.DB.Delete(&book, params["id"])
	json.NewEncoder(w).Encode(map[string]string{"status": "delete"})
}
