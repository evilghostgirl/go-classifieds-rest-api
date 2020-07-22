package handlers

import (
	"classifieds-rest-api/packages/models"
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllOffers(w http.ResponseWriter, r *http.Request) {

	classifieds, err := persistence.GetAllClassifieds()
	toJSON(w, classifieds, err)

}

func GetOfferByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	classifieds, err := persistence.GetClassifiedByID(vars["ID"])
	toJSON(w, classifieds, err)

}

func GetOffersByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	classifieds, err := persistence.GetClassifiedsByTitle(vars["title"])
	toJSON(w, classifieds, err)

}

func GetOffersByTitle1(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	fmt.Printf("title: %s", title)
	classifieds, err := persistence.GetClassifiedsByTitle(title)
	toJSON(w, classifieds, err)

}

func GetOffersByLocalizationID(w http.ResponseWriter, r *http.Request) {
	localizationID := r.FormValue("localizationid")
	classifieds, err := persistence.GetClassifiedsByLocalizationID(localizationID)
	toJSON(w, classifieds, err)

}

func toJSON(w http.ResponseWriter, classifieds []*models.Classified, err error) {
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(classifieds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
