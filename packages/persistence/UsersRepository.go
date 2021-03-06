package persistence

import (
	"classifieds-rest-api/packages/models"
	"database/sql"
)

func GetAllUsers() ([]*models.User, error) {

	rows, err := db.Query("SELECT users.id,users.first_name, users.last_name, users.localizationid, localizations.name FROM users INNER JOIN localizations ON users.localizationid = localizations.id")
	return parseToTableOfUsers(rows, err)
}

func GetUserByID(ID string) ([]*models.User, error) {
	// fmt.Printf("SELECT * FROM Users WHERE title LIKE '%s'", title)
	rows, err := db.Query("SELECT users.id,users.first_name, users.last_name, users.localizationid, localizations.name FROM users INNER JOIN localizations ON users.localizationid = localizations.id WHERE users.id = $1 ", ID)
	return parseToTableOfUsers(rows, err)
}

func GetUsersByLocalizationID(localizationID string) ([]*models.User, error) {

	rows, err := db.Query("SELECT users.id,users.first_name, users.last_name, users.localizationid, localizations.name FROM users INNER JOIN localizations ON users.localizationid = localizations.id WHERE users.localizationid = $1 ", localizationID)
	return parseToTableOfUsers(rows, err)
}

// Creates array of Users from table row data
func parseToTableOfUsers(rows *sql.Rows, err error) ([]*models.User, error) {

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		bk := new(models.User)
		err := rows.Scan(
			&bk.ID,
			&bk.FirstName,
			&bk.LastName,
			&bk.LocalizationID,
			&bk.Localization,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
