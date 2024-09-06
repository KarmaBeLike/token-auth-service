package main

import (
	"database/sql"
	"log"

	"github.com/KarmaBeLike/token-auth-service/config"
	"github.com/KarmaBeLike/token-auth-service/internal/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	runMigrations(db)
}
