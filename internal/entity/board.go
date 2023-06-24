package entity

import "time"

type BoardType string

const (
	Personal BoardType = "personal"
	Group    BoardType = "group"
)

type Board struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Type        BoardType
	ID          string
	OwnerID     string
	Title       string
	Description string
	Members     []string
	Notes       []string
}
