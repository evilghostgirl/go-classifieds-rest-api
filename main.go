package main

import (
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	persistence.InitDB("postgres://postgres:secret@localhost/postgres?sslmode=disable")

	// http.HandleFunc("/classifieds", renderAllClassifieds)
	// http.ListenAndServe(":3000", nil)

	r := mux.NewRouter()
	r.HandleFunc("/classifieds", listClassifiedsByTitle1).Methods("GET").Queries("title", "{title}")

	r.HandleFunc("/classifieds", listAllClassifieds).Methods("GET")
	r.HandleFunc("/classifieds/{title}", listClassifiedsByTitle).Methods("GET")

	http.ListenAndServe(":3000", r)

}

func listAllClassifieds(w http.ResponseWriter, r *http.Request) {

	bks, err := persistence.GetAllClassifieds()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(bks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func listClassifiedsByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bks, err := persistence.GetClassifiedsByTitle(vars["title"])
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(bks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func listClassifiedsByTitle1(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	fmt.Printf("title: %s", title)
	bks, err := persistence.GetClassifiedsByTitle(title)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(bks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
