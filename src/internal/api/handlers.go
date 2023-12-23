package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restfulGoCart/src/internal/model"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/items", createItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", deleteItem).Methods("DELETE")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.GetAllItems())
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	model.AddItem(item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item model.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	updatedItem, err := model.UpdateItem(params["id"], item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := model.DeleteItem(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("Item deleted")
}
