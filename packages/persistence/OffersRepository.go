package persistence

import (
	"classifieds-rest-api/packages/models"
	"database/sql"
	"fmt"
)

func GetAllOffers() ([]*models.Offer, error) {
	rows, err := db.Query("SELECT * FROM classifieds")
	return parseToOffers(rows, err)
}

func GetOfferByID(ID string) (*models.Offer, error) {
	var offer models.Offer
	err := db.QueryRow("SELECT * FROM classifieds WHERE id = $1 LIMIT 1", ID).Scan(&offer.ID,
		&offer.Title,
		&offer.Description,
		&offer.AddedTime,
		&offer.EndTime,
		&offer.Price,
		&offer.LocalizationID,
		&offer.OwnerID,
		&offer.CategoryID,
	)
	if err != nil {
		return &(models.Offer{}), err
	} else {
		return &offer, nil
	}
}

func GetOffersByTitle(title string) ([]*models.Offer, error) {
	rows, err := db.Query("SELECT * FROM classifieds WHERE title LIKE '%' || $1 || '%' ", title)
	return parseToOffers(rows, err)
}

func GetOffersByLocalizationID(localizationID string) ([]*models.Offer, error) {
	rows, err := db.Query("SELECT * FROM classifieds WHERE localizationid LIKE '%' || $1 || '%' ", localizationID)
	return parseToOffers(rows, err)
}

func CreateOffer(offer *models.Offer) *string {
	sqlStatement := `INSERT INTO classifieds(
		id, title, description, 
		addedtime, endtime, price, 
		localizationid, ownerid, categoryid) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING id;`
	fmt.Printf("bazka")
	var id string
	var err = db.QueryRow(sqlStatement,
		&offer.ID,
		&offer.Title,
		&offer.Description,
		&offer.AddedTime,
		&offer.EndTime,
		&offer.Price,
		&offer.LocalizationID,
		&offer.OwnerID,
		&offer.CategoryID).Scan(&id)

	if err != nil {
		id = "null"
	}
	fmt.Println("New record ID is:", id)
	return &id
}

// Creates array of Classifieds from table row data
func parseToOffers(rows *sql.Rows, err error) ([]*models.Offer, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	offers := make([]*models.Offer, 0)
	for rows.Next() {
		bk := new(models.Offer)
		err := rows.Scan(
			&bk.ID,
			&bk.Title,
			&bk.Description,
			// ,
			&bk.AddedTime,
			&bk.EndTime, &bk.Price,
			&bk.LocalizationID, &bk.OwnerID,
			&bk.CategoryID,
		)
		if err != nil {
			return nil, err
		}
		offers = append(offers, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return offers, nil
}
