package tests

import (
	"bytes"
	"classifieds-rest-api/packages/models"
	"classifieds-rest-api/packages/persistence"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestGettingCategoryByID(t *testing.T) {
	persistence.InitDB("postgres://postgres:secret@localhost/postgres?sslmode=disable")

	got, err0 := persistence.GetCategoryByID("23340f27-fb9b-43e1-b624-662c6d411a99")
	id, err := uuid.Parse("23340f27-fb9b-43e1-b624-662c6d411a99")

	if err0 != nil {
		log.Fatal("Error with getting from repository", err)
	}
	if err != nil {
		log.Fatal("Error with parsing id", err)
	}

	expected := &models.Category{
		ID:          id,
		Title:       "Ubrania",
		Description: "ciuszki",
	}

	expArray := make([]*models.Category, 0)

	expArray1 := append(expArray, expected)

	if !cmp.Equal(got, expArray1) {
		t.Errorf("GetCategoryByID('23340f27-fb9b-43e1-b624-662c6d411a99') = %+v; expected: %+v", expArray1, got)
	}
}

func TestCreateAndGetNewOffer(t *testing.T) {
	persistence.InitDB("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	var endpoint = "http://localhost:3000/offers/"
	var title = "KOSZULKA"
	offerJson := []byte(`{
		"title":"` + title + `",
		"description":"Śliczna koszulka",
		"added_time":"2004-10-19T10:23:54Z",
		"end_time":"2010-10-19T10:23:54Z",
		"price":300,
		"localization_id":"23330f27-fb9b-43e1-b624-662c5d416e99",
		"owner_id":"23340f27-fb9b-43e1-b624-662c6d411e99",
		"category_id":"23340f27-fb9b-43e1-b624-662c6d411a99"
	}
	`)

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(offerJson))

	if err != nil {
		log.Fatalln("Cannot send post request")
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	var id = result["id"]

	resp2, err2 := http.Get(endpoint + id)
	if err2 != nil {
		log.Fatalln("cannot get offer")
	}

	defer resp2.Body.Close()
	var offer models.Offer
	err3 := json.NewDecoder(resp2.Body).Decode(&offer)
	t.Logf(offer.Title)
	if err3 != nil {
		log.Fatalln("cannot decode")
	}

	if !cmp.Equal(offer.Title, title) {
		t.Errorf("offer.title != title, %s != %s", offer.Title, title)
	}

}
