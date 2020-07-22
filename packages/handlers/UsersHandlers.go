package handlers

import (
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := persistence.GetAllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	users, err := persistence.GetUserByID(vars["ID"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func GetUsersByLocalizationID(w http.ResponseWriter, r *http.Request) {
	localizationID := r.FormValue("localizationid")
	users, err := persistence.GetUsersByLocalizationID(localizationID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
