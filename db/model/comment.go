package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Comment provides comment body
type Comment struct {
	// Body provides body for comment
	Body string `json:"body"`
	// Created provides creating time for comment
	CreatedAt time.Time `json:"createdAt"`
	// AuthorID provides author of comment
	AuthorID uuid.UUID `json:"id"`
	// mentions provides definitions of mentions for comment
	Mentions []uuid.UUID `json:"mentions"`
}
