package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
)

func deleteTodo(pool *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		isDeleted := db.DeleteTodo(r.Context(), pool, id)
		if !isDeleted {
			return
		}

		tmpl.ExecuteTemplate(w, "notification", struct {
			Message string
		}{
			Message: fmt.Sprintf("Todo with id %d deleted", id),
		})
	}
}
