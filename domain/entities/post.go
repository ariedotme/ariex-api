package entities

import "time"

type Post struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	NormalizedTitle string    `json:"normalized_title"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
