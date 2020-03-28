package main

import (
	"classifieds-rest-api/packages/handlers"
	"classifieds-rest-api/packages/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	persistence.InitDB("postgres://postgres:secret@localhost/postgres?sslmode=disable")

	// http.HandleFunc("/classifieds", renderAllClassifieds)
	// http.ListenAndServe(":3000", nil)

	r := mux.NewRouter()
	r.HandleFunc("/classifieds", handlers.ListClassifiedsByTitle1).Methods("GET").Queries("title", "{title}")
	r.HandleFunc("/classifieds", handlers.ListClassifiedsByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	r.HandleFunc("/classifieds", handlers.ListAllClassifieds).Methods("GET")
	r.HandleFunc("/classifieds/", handlers.ListAllClassifieds).Methods("GET")
	r.HandleFunc("/classifieds/{ID}", handlers.ListClassifiedByID).Methods("GET")

	r.HandleFunc("/categories/{ID}", handlers.ListCategoryByID).Methods("GET")
	r.HandleFunc("/categories", handlers.ListAllCategories).Methods("GET")
	r.HandleFunc("/categories/", handlers.ListAllCategories).Methods("GET")

	r.HandleFunc("/users", handlers.ListUsersByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	r.HandleFunc("/users/{ID}", handlers.ListUserByID).Methods("GET")
	r.HandleFunc("/users", handlers.ListAllUsers).Methods("GET")
	r.HandleFunc("/users/", handlers.ListAllUsers).Methods("GET")

	http.ListenAndServe(":3000", r)

}
