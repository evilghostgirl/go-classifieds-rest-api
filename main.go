package main

import (
	"classifieds-rest-api/packages/config"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", config.CheckHealth).Methods("GET")
	r.HandleFunc("/register", config.GetConsulRegisterRequest).Methods("GET")

	config.RegisterInServiceDiscovery()

	// persistence.InitDB("postgres://hazelina:secret@db/postgres?sslmode=disable")
	// r.HandleFunc("/classifieds", handlers.ListClassifiedsByTitle1).Methods("GET").Queries("title", "{title}")
	// r.HandleFunc("/classifieds", handlers.ListClassifiedsByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	// r.HandleFunc("/classifieds", handlers.ListAllClassifieds).Methods("GET")
	// r.HandleFunc("/classifieds/", handlers.ListAllClassifieds).Methods("GET")
	// r.HandleFunc("/classifieds/{ID}", handlers.ListClassifiedByID).Methods("GET")

	// r.HandleFunc("/categories/{ID}", handlers.ListCategoryByID).Methods("GET")
	// r.HandleFunc("/categories", handlers.ListAllCategories).Methods("GET")
	// r.HandleFunc("/categories/", handlers.ListAllCategories).Methods("GET")

	// r.HandleFunc("/users", handlers.ListUsersByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	// r.HandleFunc("/users/{ID}", handlers.ListUserByID).Methods("GET")
	// r.HandleFunc("/users", handlers.ListAllUsers).Methods("GET")
	// r.HandleFunc("/users/", handlers.ListAllUsers).Methods("GET")

	// http.HandleFunc("/classifieds", renderAllClassifieds)
	// http.ListenAndServe(":3000", nil)

	http.ListenAndServe(":3000", r)

}
