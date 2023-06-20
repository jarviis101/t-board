package main

import (
	"context"
	"log"
	"t-mail/internal/app"
	"t-mail/pkg"
)

func main() {
	config, err := pkg.CreateConfig()
	if err != nil {
		log.Panic(err.Error())
	}
	db := pkg.CreateDatabaseConnection(context.Background(), config.Database)

	application := app.CreateApplication(db)
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
