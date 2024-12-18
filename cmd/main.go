package main

import (
	"database/sql"
	_ "github.com/justyork/api-template/docs" // Import generated docs
	"github.com/justyork/api-template/internal/middleware"
	"github.com/swaggo/http-swagger" // Swagger handler
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/justyork/api-template/internal/routes"
	_ "github.com/mattn/go-sqlite3"
)

func applyMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}

// @title API Template
// @version 1.0
// @description A lightweight and scalable REST API server built with GoLang.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Get environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}
	middleware.SetJWTKey([]byte(jwtSecret))

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	applyMigrations(databaseURL)

	// Register routes
	r := routes.RegisterRoutes()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Printf("Starting server on :%s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
