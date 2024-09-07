package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/KarmaBeLike/token-auth-service/config"
	"github.com/KarmaBeLike/token-auth-service/internal/database"
	"github.com/KarmaBeLike/token-auth-service/internal/handlers"
	"github.com/KarmaBeLike/token-auth-service/internal/repository"
	"github.com/KarmaBeLike/token-auth-service/internal/service"

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

	userRepo := &repository.UserRepo{DB: db}
	userService := &service.UserService{Repository: *userRepo}

	// Создание обработчиков
	userHandler := &handlers.UserHandler{UserService: userService}

	// Маршруты
	http.HandleFunc("/register", userHandler.RegisterUser)
	http.HandleFunc("/login", userHandler.LoginUser)

	// Запуск сервера
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
