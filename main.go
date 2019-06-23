package main 

import (
"log"
"net/http"
"Book-List/models"
"Book-List/driver"
"Book-List/controllers"
"database/sql"
"github.com/gorilla/mux"
"github.com/subosito/gotenv"
)


var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
router := mux.NewRouter()
db = driver.ConnectDB()
controller := controllers.Controller{}

router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

log.Fatal(http.ListenAndServe(":8000", router))
}