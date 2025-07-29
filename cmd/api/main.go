package main

import (
	"log"

	"github.com/FarzadMohtasham/go-social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server has started at %s", app.config.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}
