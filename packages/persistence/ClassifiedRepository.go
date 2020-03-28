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

func GetClassifiedByID(ID string) ([]*models.Classified, error) {
	// fmt.Printf("SELECT * FROM classifieds WHERE title LIKE '%s'", title)
	rows, err := db.Query("SELECT * FROM classifieds WHERE id = $1", ID)
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

func GetClassifiedsByTitle(title string) ([]*models.Classified, error) {
	// fmt.Printf("SELECT * FROM classifieds WHERE title LIKE '%s'", title)
	rows, err := db.Query("SELECT * FROM classifieds WHERE title LIKE '%' || $1 || '%' ", title)
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

func GetClassifiedsByLocalizationID(localizationID string) ([]*models.Classified, error) {
	// fmt.Printf("SELECT * FROM classifieds WHERE title LIKE '%s'", title)
	rows, err := db.Query("SELECT * FROM classifieds WHERE localizationid LIKE '%' || $1 || '%' ", localizationID)
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
