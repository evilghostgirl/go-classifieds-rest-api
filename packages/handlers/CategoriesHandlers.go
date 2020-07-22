package handlers

import (
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := persistence.GetAllCategories()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	categories, err := persistence.GetCategoryByID(vars["ID"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
