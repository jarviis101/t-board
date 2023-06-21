package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Boards    []primitive.ObjectID `bson:"boards,omitempty"`
	CreatedAt time.Time            `bson:"created_at,omitempty"`
	UpdatedAt time.Time            `bson:"updated_at,omitempty"`
	Name      string               `bson:"name,omitempty"`
	Email     string               `bson:"email,omitempty"`
	Password  string               `bson:"password,omitempty"`
}

type Note struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Board       primitive.ObjectID `bson:"board,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
	Description string             `bson:"description,omitempty"`
}

type Board struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Members     []primitive.ObjectID `bson:"members,omitempty"`
	CreatedAt   time.Time            `bson:"created_at,omitempty"`
	UpdatedAt   time.Time            `bson:"updated_at,omitempty"`
	Title       string               `bson:"title,omitempty"`
	Description string               `bson:"description,omitempty"`
}
