package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olemart1n/nub.global/db"
)

func test(pool *pgxpool.Pool, tmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update of todo")

		task := r.FormValue("task")

		var done int
		doneInput := r.FormValue("done")
		if doneInput == "on" {
			done = 1
		} else {
			done = 0
		}
		id, _ := strconv.Atoi(r.FormValue("id"))

		if task == "" {
			http.Error(w, "Missing task OR done OR id", http.StatusBadRequest)
			return
		}

		todo := db.Todo{Task: task, Done: done, ID: id}
		insertedTodo, err := db.AddTodo(r.Context(), pool, todo)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "todo", insertedTodo)
	}

}
