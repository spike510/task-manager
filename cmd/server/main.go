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

// @title Task Manager API
// @version 1.0
// @description API do zarządzania zadaniami (TODO).
// @host localhost:8080
// @BasePath /
func main() {
	// Database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	// Run migrations
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatal("Database instance init failed:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Migrate instance init failed:", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal("Migrations failed:", err)
	}

	log.Println("Migrations applied successfully")

	r := http.NewRouter(database)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on :8080")
}
