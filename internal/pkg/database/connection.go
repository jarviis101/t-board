package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type database struct {
	client *mongo.Client
}

func CreateDatabaseConnection(ctx context.Context) *mongo.Database {
	db := database{}
	return db.openConnection(ctx)
}

func (d *database) openConnection(ctx context.Context) *mongo.Database {
	d.resolveClient(ctx)
	return d.client.Database("t-mail")
}

func (d *database) resolveClient(ctx context.Context) {
	if d.client != nil {
		return
	}
	opts := options.Client().ApplyURI("mongodb://user:password@localhost:27017/t-mail")
	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	c.Ping(ctx, nil)

	d.client = c
}
