package main

import (
	"log"
	"t-board/internal/app/ws"
)

func main() {
	//config, err := pkg.CreateConfig()
	//if err != nil {
	//	log.Panic(err.Error())
	//}
	//db := pkg.CreateDatabaseConnection(context.Background(), config.Database)
	//
	//container := di.CreateContainer(db, config.Server)
	application := ws.CreateApplication()
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
