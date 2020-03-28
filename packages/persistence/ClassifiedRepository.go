package persistence

import (
	"classifieds-rest-api/packages/models"
	"database/sql"
)

func GetAllClassifieds() ([]*models.Classified, error) {
	rows, err := db.Query("SELECT * FROM classifieds")
	return parseToTableOfClassifieds(rows, err)
}

func GetClassifiedByID(ID string) ([]*models.Classified, error) {
	rows, err := db.Query("SELECT * FROM classifieds WHERE id = $1", ID)
	return parseToTableOfClassifieds(rows, err)
}

func GetClassifiedsByTitle(title string) ([]*models.Classified, error) {
	rows, err := db.Query("SELECT * FROM classifieds WHERE title LIKE '%' || $1 || '%' ", title)
	return parseToTableOfClassifieds(rows, err)
}

func GetClassifiedsByLocalizationID(localizationID string) ([]*models.Classified, error) {
	rows, err := db.Query("SELECT * FROM classifieds WHERE localizationid LIKE '%' || $1 || '%' ", localizationID)
	return parseToTableOfClassifieds(rows, err)
}

// Creates array of Classifieds from table row data
func parseToTableOfClassifieds(rows *sql.Rows, err error) ([]*models.Classified, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classifieds := make([]*models.Classified, 0)
	for rows.Next() {
		bk := new(models.Classified)
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
		classifieds = append(classifieds, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return classifieds, nil
}
