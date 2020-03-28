package persistence

import (
	"classifieds-rest-api/packages/models"
)

func GetAllCategories() ([]*models.Category, error) {
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*models.Category, 0)
	for rows.Next() {
		bk := new(models.Category)
		err := rows.Scan(
			&bk.ID,
			&bk.Title,
			&bk.Description,
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

func GetCategoryByID(ID string) ([]*models.Category, error) {
	rows, err := db.Query("SELECT * FROM categories WHERE id = $1 ", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*models.Category, 0)
	for rows.Next() {
		bk := new(models.Category)
		err := rows.Scan(
			&bk.ID,
			&bk.Title,
			&bk.Description,
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
