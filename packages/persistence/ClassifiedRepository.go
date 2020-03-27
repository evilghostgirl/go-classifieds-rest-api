package persistence

import (
	"classifieds-rest-api/packages/models"
)

func GetAllClassifieds() ([]*models.Classified, error) {
	rows, err := db.Query("SELECT * FROM classifieds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*models.Classified, 0)
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
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func GetClassifiedsByTitle(title string) ([]*models.Classified, error) {
	// fmt.Printf("SELECT * FROM classifieds WHERE title LIKE '%s'", title)
	rows, err := db.Query("SELECT * FROM classifieds WHERE title LIKE '%' || $1 || '%' ", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*models.Classified, 0)
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
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
