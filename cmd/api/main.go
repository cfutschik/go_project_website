package main

import (
	"log"

	"github.com/cfutschik/go_project_website.git/internal/db"
	"github.com/cfutschik/go_project_website.git/internal/env"
	"github.com/cfutschik/go_project_website.git/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":4001"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://christianfutschik:adminpassword@localhost/myprojectdb?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panicln("new db: %w", err)
	}
	log.Printf("db connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
