package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetTodos(ctx context.Context, pool *pgxpool.Pool) ([]Todo, error) {
	query := "SELECT id, task, done FROM todos"
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todos []Todo

	for rows.Next() {
		var todo Todo
		rowErr := rows.Scan(&todo.ID, &todo.Task, &todo.Done)
		if rowErr != nil {
			return nil, rowErr
		}
		todos = append(todos, todo)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
