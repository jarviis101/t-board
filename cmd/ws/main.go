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

	container := di.CreateContainer(db, config.Server)
	application := ws.CreateApplication(container)
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
