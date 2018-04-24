package model

import "github.com/satori/go.uuid"

// Author provides definition for authour struct
type Author struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	JoinedAt    string    `json:"joinedAt"`
}
