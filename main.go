package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Smoothie struct untuk respont json
type Smoothie struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Ingredients []string `json:"inggredients"`
}

// Smoothies adalah kumpulan dari smoothie
var Smoothies []Smoothie

func allSmoothie(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Smoothies)
}
func showSmoothie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, smoothie := range Smoothies {
		if smoothie.ID == key {
			json.NewEncoder(w).Encode(smoothie)
		}
	}
}
func addSmoothie(w http.ResponseWriter, r *http.Request) {
	// Mengambil body dari POST
	// Unmarshall ke Smoothie struct
	// Tambahkan ke array Smoothies
	body, _ := ioutil.ReadAll(r.Body)
	var smoothie Smoothie

	json.Unmarshal(body, &smoothie)
	// update Array Smoothies
	Smoothies = append(Smoothies, smoothie)
	json.NewEncoder(w).Encode(smoothie)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jualan Minooman")
}
func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/smoothies", allSmoothie).Methods("GET")
	router.HandleFunc("/smoothie", addSmoothie).Methods("POST")
	router.HandleFunc("/smoothie/{id}", showSmoothie).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Smoothies = []Smoothie{
		{ID: "1", Name: "Es Jeruk", Ingredients: []string{"Air dingin", "Jeruk"}},
		{ID: "2", Name: "Es Teh", Ingredients: []string{"Air dingin", "Teh"}},
	}
	handleRequest()
}
