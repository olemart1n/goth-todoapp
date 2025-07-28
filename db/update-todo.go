package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateTodo(ctx context.Context, pool *pgxpool.Pool, todo Todo) (Todo, error) {
	var insertedTodo Todo
	query := "UPDATE todos set task = $1, done = $2 WHERE id = $3 RETURNING id, task, done"
	err := pool.QueryRow(ctx, query, todo.Task, todo.Done).Scan(&insertedTodo.ID, &insertedTodo.Task, &insertedTodo.Done)
	if err != nil {
		return Todo{}, fmt.Errorf("failed to insert todo: %w", err)
	}

	return insertedTodo, nil
}
