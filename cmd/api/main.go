package main

import "log"

func main() {
	cfg := config{
		addr: ":8880",
	}

	app := &application{
		config: cfg,
	}

	log.Printf("Server has started at %s", app.config.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}
