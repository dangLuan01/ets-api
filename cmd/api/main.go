package main

import (
	"log"

	"github.com/dangLuan01/ets-api/internal/app"
	"github.com/dangLuan01/ets-api/internal/config"
)

func main() {
	
	app.LoadEnv()
	
	cfg := config.NewConfig()
	
	application, err := app.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to start app:%s", err)
	}

	if err := application.Run(); err != nil {
		panic(err)
	}
}