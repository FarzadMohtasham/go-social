package main

import (
	"log"

	"github.com/FarzadMohtasham/go-social/internal/env"
	"github.com/FarzadMohtasham/go-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Printf("Server has started at %s", app.config.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}
