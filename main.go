package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Smoothie struct untuk respont json
type Smoothie struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"inggredients"`
}

// Smoothies adalah kumpulan dari smoothie
var Smoothies []Smoothie

func returnAllSmothies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Smoothies)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jualan Minooman")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/minooman", returnAllSmothies)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	Smoothies = []Smoothie{
		{Name: "Es Jeruk", Ingredients: []string{"Air dingin", "Jeruk"}},
		{Name: "Es Teh", Ingredients: []string{"Air dingin", "Teh"}},
	}
	handleRequest()
}