package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

var books []Book

// Get all books
func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

// Get one book
func getBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Book{})
}

// Create book
func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}

// Update book
func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	for index, item := range books {
		if item.ID == params["id"] {
			books[index].Title = book.Title
			books[index].Isbn = book.Isbn
			books[index].Author = book.Author
			json.NewEncoder(res).Encode(books)
			return
		}
	}
}

// Delete book
func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var book Book
	for index, item := range books {
		if item.ID == params["id"] {
			book = item
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(book)
}

func main() {
	// Create mock data
	books = append(books, Book{ID: "1", Isbn: "354231", Title: "Book one", Author: &Author{Firstname: "Daniel", Lastname: "Yokoyama"}})
	books = append(books, Book{ID: "2", Isbn: "367231", Title: "Book two", Author: &Author{Firstname: "Lohana", Lastname: "Soares"}})

	// Init router
	router := mux.NewRouter()

	// Create route handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	log.Fatal(http.ListenAndServe(":8000", router))
}
