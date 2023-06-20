package main

import (
	"log"
	"t-mail/internal/app"
)

func main() {
	application := app.CreateApplication()
	if err := application.Run(); err != nil {
		log.Panic(err.Error())
	}
}
