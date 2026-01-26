package main

import (
	"log"
	"wallet-app/config"
	"wallet-app/internal/db"
	migrate "wallet-app/migrations"

	_ "github.com/lib/pq"
)

func main() {
	dsn, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.ConnectDB(dsn.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrate.Run(database.DB); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully")
}
