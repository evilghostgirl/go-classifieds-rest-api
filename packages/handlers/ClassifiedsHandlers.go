package handlers

import (
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ListAllClassifieds(w http.ResponseWriter, r *http.Request) {

	classifieds, err := persistence.GetAllClassifieds()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ListClassifiedByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	classifieds, err := persistence.GetClassifiedByID(vars["ID"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ListClassifiedsByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	classifieds, err := persistence.GetClassifiedsByTitle(vars["title"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ListClassifiedsByTitle1(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	fmt.Printf("title: %s", title)
	classifieds, err := persistence.GetClassifiedsByTitle(title)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ListClassifiedsByLocalizationID(w http.ResponseWriter, r *http.Request) {
	localizationID := r.FormValue("localizationid")
	classifieds, err := persistence.GetClassifiedsByLocalizationID(localizationID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
