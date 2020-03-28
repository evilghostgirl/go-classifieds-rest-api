package persistence

import (
	"classifieds-rest-api/packages/models"
	"database/sql"
)

func GetAllCategories() ([]*models.Category, error) {
	rows, err := db.Query("SELECT * FROM categories")
	return parseToTableOfCategories(rows, err)
}

func GetCategoryByID(ID string) ([]*models.Category, error) {
	rows, err := db.Query("SELECT * FROM categories WHERE id = $1 ", ID)
	return parseToTableOfCategories(rows, err)
}

// Creates array of Categories from table row data
func parseToTableOfCategories(rows *sql.Rows, err error) ([]*models.Category, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*models.Category, 0)
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
		categories = append(categories, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
