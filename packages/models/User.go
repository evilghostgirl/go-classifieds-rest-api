package models

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	LocalizationID uuid.UUID `json:"localization_id"`
	Localization   string    `json:"localization"`
}
