package routes

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
)

func GetUpdateForm(pool *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var todo db.Todo
		id, _ := strconv.Atoi(vars["id"])
		todo, err := db.GetTodoWithID(r.Context(), pool, id)
		if err != nil {
			return
		}
		tmpl.ExecuteTemplate(w, "update-form", todo)
	}
}
