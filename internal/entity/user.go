package entity

import "time"

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        string
	Name      string
	Email     string
	Password  string
	Boards    []string
}
