package model

import "github.com/satori/go.uuid"

// Author provides definition for authour struct
type Author struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}