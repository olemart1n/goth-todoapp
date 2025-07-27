package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
)

func newTodo(pool *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New todo incoming")

		task := r.FormValue("task")
		if task == "" {
			http.Error(w, "Missing task", http.StatusBadRequest)
			return
		}

		todo := db.Todo{Task: task, Done: 0}
		insertedTodo, err := db.AddTodo(r.Context(), pool, todo)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "todo", insertedTodo)
	}

}
