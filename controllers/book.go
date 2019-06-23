package controllers

import (
"log"
"net/http"
"strconv"
"Book-List/models"
"encoding/json"
"github.com/gorilla/mux"
"database/sql"
"Book-List/repository/book"
)

type Controller struct {}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}
	repository := bookRepository.BookRepository{}
	books = repository.GetBooks(db, book, books)
	json.NewEncoder(w).Encode(books)

}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	repository := bookRepository.BookRepository{}
	book = repository.GetBook(db, book, id)
	
	json.NewEncoder(w).Encode(book)
}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	repository := bookRepository.BookRepository {}
	bookID = repository.AddBook(db, book)

	json.NewEncoder(w).Encode(bookID)
}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	repository := bookRepository.BookRepository {}
	rowsUpdated := repository.UpdateBook(db, book)

    json.NewEncoder(w).Encode(rowsUpdated)
}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	logFatal(err)
	
	repository := bookRepository.BookRepository {}
	rowsDeleted := repository.RemoveBook(db, id)

	json.NewEncoder(w).Encode(rowsDeleted)
}
}