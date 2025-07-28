package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool, tmpl *template.Template) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", IndexHandler(pool, tmpl)).Methods("GET")
	r.HandleFunc("/get-update-form/{id}", GetUpdateForm(pool, tmpl)).Methods("GET")
	r.HandleFunc("/new-todo", newTodo(pool, tmpl)).Methods("POST")

	r.HandleFunc("/update-todo/{id}", updateTodo(pool, tmpl)).Methods("PUT")
	r.HandleFunc("/test", test(pool, tmpl)).Methods("PUT")

	r.HandleFunc("/todo/{id}", deleteTodo(pool, tmpl)).Methods("DELETE")

	// Serve static files at /static/
	staticFileDirectory := http.Dir("static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	// etc...
	return r
}
