package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type database struct {
	client *mongo.Client
}

func CreateDatabaseConnection(ctx context.Context, cfg Database) *mongo.Database {
	db := database{}
	return db.openConnection(ctx, cfg)
}

func (d *database) openConnection(ctx context.Context, cfg Database) *mongo.Database {
	d.resolveClient(ctx, cfg)
	return d.client.Database(cfg.Name)
}

func (d *database) resolveClient(ctx context.Context, cfg Database) {
	if d.client != nil {
		return
	}
	opts := options.Client().ApplyURI(cfg.Uri)
	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatalln(err.Error())
	}

	d.client = c
}
