package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func applyMigrations(db *sql.DB) {
	m, err := migrate.New(
		"file://migrations",
		"sqlite3://db.sqlite",
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	applyMigrations(db)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running"))
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
