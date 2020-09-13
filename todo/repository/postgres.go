package repository

import (
	"context"
	"database/sql"

	"github.com/ks6088ts/spidey/todo/model"

	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) PutTodo(ctx context.Context, a model.Todo) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO todos(id, name) VALUES($1, $2)", a.ID, a.Name)
	return err
}

func (r *postgresRepository) GetTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM todos WHERE id = $1", id)
	a := &model.Todo{}
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *postgresRepository) ListTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name FROM todos ORDER BY id DESC OFFSET $1 LIMIT $2",
		skip,
		take,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		a := &model.Todo{}
		if err = rows.Scan(&a.ID, &a.Name); err == nil {
			todos = append(todos, *a)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
