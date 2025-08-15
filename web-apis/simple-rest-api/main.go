package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items []Item
var nextID = 1

func main() {
	// Initialize with some sample data
	items = append(items, Item{ID: nextID, Name: "Laptop", Price: 1000})
	nextID++
	items = append(items, Item{ID: nextID, Name: "Mouse", Price: 25})
	nextID++

	// Create router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start server
	fmt.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = nextID
	nextID++
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range items {
		if item.ID == id {
			items[i].Name = r.FormValue("name")
			items[i].Price, _ = strconv.Atoi(r.FormValue("price"))
			json.NewEncoder(w).Encode(items[i])
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Item deleted"})
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
