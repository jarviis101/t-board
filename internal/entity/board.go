package entity

import "time"

type Board struct {
	ID          string
	Title       string
	Description string
	Members     []string
	Notes       []string
	CreatedAt   time.Time
}
