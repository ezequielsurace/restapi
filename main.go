package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// INIT BOOKS VAR AS A SLICE BOOK STRUCT
var books []Book

// GET ALL BOOKS
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET SINGLE BOOK
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

// CREATE BOOK
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	var _ = json.NewDecoder(r.Body).Decode(&book)
	lastID, err := strconv.Atoi(books[len(books)-1].ID)
	if err != nil {
		log.Fatalln("Error createBook", err)
	}
	book.ID = strconv.Itoa(lastID + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// UPDATE BOOK
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books[index].Isbn = book.Isbn
			books[index].Title = book.Title
			books[index].Author.Firstname = book.Author.Firstname
			books[index].Author.Lastname = book.Author.Lastname
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}

}

// DELETE BOOK
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {

	// Init Router
	r := mux.NewRouter()

	// MOCK DATA
	books = append(books, Book{ID: "1", Isbn: "0747532699", Title: "Harry Potter and the Philosopher's Stone", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "2", Isbn: "0747538492", Title: "Harry Potter and the Chamber of Secrets", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "3", Isbn: "0747542155", Title: "Harry Potter and the Prisoner of Azkaban", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "4", Isbn: "074754624X", Title: "Harry Potter and the Goblet of Fire", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "5", Isbn: "0747551006", Title: "Harry Potter and the Order of the Phoenix", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "6", Isbn: "0747581088", Title: "Harry Potter and the Half-Blood Prince", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "7", Isbn: "0545010225", Title: "Harry Potter and the Deathly Hallows", Author: &Author{Firstname: "J. K.", Lastname: "Rowling"}})
	books = append(books, Book{ID: "8", Isbn: "9780385737944", Title: "The Maze Runner", Author: &Author{Firstname: "James", Lastname: "Dashner"}})
	books = append(books, Book{ID: "9", Isbn: "9780385738750", Title: "The Scorch Trials", Author: &Author{Firstname: "James", Lastname: "Dashner"}})
	books = append(books, Book{ID: "10", Isbn: "9780385738774", Title: "The Death Cure", Author: &Author{Firstname: "James", Lastname: "Dashner"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
