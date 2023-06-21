package entity

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Boards    []string
	CreatedAt time.Time
}
