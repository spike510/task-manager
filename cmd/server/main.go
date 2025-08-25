package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spike510/task-manager/internal/db"
	"github.com/spike510/task-manager/internal/http"
)

func main() {
	// Database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer database.Close()

	// Run migrations
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres",
		driver,
	)
	m.Up()

	log.Println("Migrations applied successfully")

	r := http.NewRouter(database)
	if err := r.Run(":8085"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on :8085")
}
