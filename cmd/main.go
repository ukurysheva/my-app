package main

import (
	"log"

	myapp "github.com/ukurysheva/my-app"
	"github.com/ukurysheva/my-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(myapp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
