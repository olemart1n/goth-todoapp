// Package routes handles for all routes
package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
)

func IndexHandler(pool *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request to '%s' incoming (Method: %s, User-Agent: %s)\n", r.URL.Path, r.Method, r.UserAgent())

		todos, dbErr := db.GetTodos(r.Context(), pool)
		if dbErr != nil {
			http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		}
		tmpl.ExecuteTemplate(w, "index.html", todos)

	}

}
