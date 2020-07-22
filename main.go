package main

import (
	"classifieds-rest-api/packages/handlers"
	"classifieds-rest-api/packages/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/health", config.CheckHealth).Methods("GET")
	// r.HandleFunc("/register", config.GetConsulRegisterRequest).Methods("GET")

	// config.RegisterInServiceDiscovery()

	persistence.InitDB("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	r.HandleFunc("/offers", handlers.GetOffersByTitle1).Methods("GET").Queries("title", "{title}")
	r.HandleFunc("/offers", handlers.GetOffersByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	r.HandleFunc("/offers", handlers.GetAllOffers).Methods("GET")
	r.HandleFunc("/offers/", handlers.GetAllOffers).Methods("GET")
	r.HandleFunc("/offers/{ID}", handlers.GetOfferByID).Methods("GET")

	r.HandleFunc("/categories/{ID}", handlers.GetCategoryByID).Methods("GET")
	r.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	r.HandleFunc("/categories/", handlers.GetAllCategories).Methods("GET")

	r.HandleFunc("/users", handlers.GetUsersByLocalizationID).Methods("GET").Queries("localizationid", "{localizationid}")
	r.HandleFunc("/users/{ID}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/", handlers.GetAllUsers).Methods("GET")

	http.ListenAndServe(":3000", r)

}
