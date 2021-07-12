package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Name   string `json:"Name"`
	Author string `json:"Author"`
	Genre  string `json:"Genre"`
}

type Books []Book

func ListofBooks(w http.ResponseWriter, r *http.Request) {
	books := Books{
		Book{Name: "Wings of fire", Author: "APJ", Genre: "Motivational"},
	}
	fmt.Println("Endpoint Hit: All Books Endpoint")
	json.NewEncoder(w).Encode(books)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/books", ListofBooks)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
