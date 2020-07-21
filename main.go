package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

//Get A Books
func getBook(w http.ResponseWriter, r *http.Request) {

}

//Create A Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update A Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//Delete A Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

//Delete All Books
func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data

	//Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
