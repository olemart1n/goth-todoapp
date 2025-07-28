package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteTodo(ctx context.Context, pool *pgxpool.Pool, id int) bool {
	query := "DELETE FROM todos WHERE id = $1 RETURNING id"
	var deletedTodoID int
	err := pool.QueryRow(ctx, query, id).Scan(&deletedTodoID)
	if err != nil {
		fmt.Println("failed to delete todo: %w", err)

		return false
	}

	return true
}
