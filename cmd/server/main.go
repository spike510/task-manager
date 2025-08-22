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

	log.Println("migrations applied successfully")

	r := http.NewRouter()

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
