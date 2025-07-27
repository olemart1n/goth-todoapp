// package main
package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
	"github.com/olemart1n/nub.global/routes"
)

type PageData struct {
	Tasks []db.Todo
}

func main() {
	var pool *pgxpool.Pool
	dbCtx := context.Background()
	cfg, _ := LoadConfig()
	pool, err := pgxpool.New(dbCtx, cfg.DatabaseURL)
	if err != nil {
		log.Println("Error when creating a pool")
		return
	}
	defer pool.Close()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	router := routes.NewRouter(pool, tmpl)

	if err := pool.Ping(dbCtx); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL!")

	http.ListenAndServe(":8080", router)
}
