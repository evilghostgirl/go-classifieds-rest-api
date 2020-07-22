package handlers

import (
	"classifieds-rest-api/packages/models"
	"classifieds-rest-api/packages/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllOffers(w http.ResponseWriter, r *http.Request) {
	offers, err := services.GetAllOffers()
	toJSON(w, offers, err)

}

func GetOfferByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	offer, err := services.GetOfferByID(id)
	fmt.Printf("id: %s", id)
	entityToJSON(w, offer, err)
}

func GetOffersByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	offers, err := services.GetOffersByTitle(vars["title"])
	toJSON(w, offers, err)
}

func GetOffersByTitle1(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	fmt.Printf("title: %s", title)
	offers, err := services.GetOffersByTitle(title)
	toJSON(w, offers, err)
}

func GetOffersByLocalizationID(w http.ResponseWriter, r *http.Request) {
	localizationID := r.FormValue("localizationid")
	offers, err := services.GetOffersByLocalization(localizationID)
	toJSON(w, offers, err)
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var offer models.Offer
	_ = json.NewDecoder(r.Body).Decode(&offer)
	offer.ID = uuid.New()
	fmt.Printf("'%s'", offer.Title)
	var id *string
	id = services.CreateOffer(&offer)
	idDto := IdDTO{Id: id}
	json.NewEncoder(w).Encode(&idDto)
}

type IdDTO struct {
	Id *string `json:"id"`
}

func toJSON(w http.ResponseWriter, offers []*models.Offer, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}

	if err := json.NewEncoder(w).Encode(offers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func entityToJSON(w http.ResponseWriter, offer *models.Offer, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	fmt.Println("New record ID is:", offer.ID)
	fmt.Println("New record title is:", offer.Title)
	if err := json.NewEncoder(w).Encode(offer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
