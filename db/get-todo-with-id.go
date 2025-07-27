package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTodoWithID(ctx context.Context, pool *pgxpool.Pool, id int) (Todo, error) {
	query := "SELECT id, task, done FROM todos WHERE id = $1"
	row := pool.QueryRow(ctx, query, id)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Task, &todo.Done)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}
