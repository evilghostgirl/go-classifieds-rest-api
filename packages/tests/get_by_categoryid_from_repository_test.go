package tests

import (
	"classifieds-rest-api/packages/models"
	"classifieds-rest-api/packages/persistence"
	"log"
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
