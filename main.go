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
	// mengambil parameter dari vars
	vars := mux.Vars(r)
	// ambill key id
	key := vars["id"]
	// looping dan cek id yang sama dengan key
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
func updateSmoothie(w http.ResponseWriter, r *http.Request) {
	// Mengambil body dari POST
	// Unmarshall ke Smoothie struct
	// Tambahkan ke array Smoothies
	body, _ := ioutil.ReadAll(r.Body)
	var smoothie Smoothie
	json.Unmarshal(body, &smoothie)

	// mengambil parameter dari vars
	vars := mux.Vars(r)
	// ambill key id
	key := vars["id"]
	// looping array dari smoothie kemudan cek idnya
	for index, smoothie := range Smoothies {
		// jika id sama dengan key
		if smoothie.ID == key {
			// update Array Smoothies
			Smoothies = append(Smoothies[:index], smoothie)
		}
	}
}
func deleteSmoothie(w http.ResponseWriter, r *http.Request) {
	// mengambil parameter dari vars
	vars := mux.Vars(r)
	// ambill key id
	key := vars["id"]

	// looping array dari smoothie kemudan cek idnya
	for index, smoothie := range Smoothies {
		// jika id sama dengan key
		if smoothie.ID == key {
			// kemudian hapus smoothie
			Smoothies = append(Smoothies[:index], Smoothies[index+1:]...)
		}
	}
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
	router.HandleFunc("/smoothie/{id}", updateSmoothie).Methods("PUT")
	router.HandleFunc("/smoothie/{id}", deleteSmoothie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Smoothies = []Smoothie{
		{ID: "1", Name: "Es Jeruk", Ingredients: []string{"Air dingin", "Jeruk"}},
		{ID: "2", Name: "Es Teh", Ingredients: []string{"Air dingin", "Teh"}},
	}
	handleRequest()
}
