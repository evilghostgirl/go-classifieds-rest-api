package services

import (
	"classifieds-rest-api/packages/models"
	"classifieds-rest-api/packages/persistence"
	"fmt"
)

func GetAllOffers() ([]*models.Offer, error) {
	offers, err := persistence.GetAllOffers()
	return offers, err
}

func GetOfferByID(id string) (*models.Offer, error) {
	offer, err := persistence.GetOfferByID(id)
	return offer, err
}

func GetOffersByTitle(title string) ([]*models.Offer, error) {
	offers, err := persistence.GetOffersByTitle(title)
	return offers, err
}

func GetOffersByTitle1(title string) ([]*models.Offer, error) {
	offers, err := persistence.GetOffersByTitle(title)
	return offers, err
}

func GetOffersByLocalization(id string) ([]*models.Offer, error) {
	offers, err := persistence.GetOffersByLocalizationID(id)
	return offers, err
}

func CreateOffer(offer *models.Offer) *string {
	fmt.Printf("'%s'", offer.Title)
	id := persistence.CreateOffer(offer)
	return id
}
