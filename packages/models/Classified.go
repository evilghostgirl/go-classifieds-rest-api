package models

import (
	"time"

	"github.com/google/uuid"
)

type Classified struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	AddedTime      time.Time `json:"added_time"`
	EndTime        time.Time `json:"end_time"`
	Price          float64   `json:"Price"`
	LocalizationID uuid.UUID `json:"localization_id"`
	OwnerID        uuid.UUID `json:"owner_id"`
	CategoryID     uuid.UUID `json:"category_id"`
}
