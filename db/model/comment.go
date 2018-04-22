package model

import "time"

// Comment provides comment body
type Comment struct {
	// Body provides body for comment
	Body string `json:"body"`
	// Created provides creating time for comment
	CreatedAt time.Time `json:"createdAt"`
}
