package main

import (
	"log"

	"github.com/cfutschik/go_project_website.git/internal/env"
	"github.com/cfutschik/go_project_website.git/internal/env/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4001"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
