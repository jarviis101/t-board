package main

import (
	"context"
	"log"
	"t-board/internal/app/di"
	"t-board/internal/app/ws"
	"t-board/pkg"
)

func main() {
	config, err := pkg.CreateConfig()
	if err != nil {
		log.Panic(err.Error())
	}

	db := pkg.CreateDatabaseConnection(context.Background(), config.Database)
	defer func() {
		if err := pkg.CloseConnection(db.Client()); err != nil {
			log.Panic(err.Error())
		}
	}()

	container := di.CreateContainer(db, config.Server)
	application := ws.CreateApplication(container)
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
