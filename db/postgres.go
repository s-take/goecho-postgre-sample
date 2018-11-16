package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/s-take/goecho-postgre-sample/schema"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertTask(ctx context.Context, task schema.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks(id, body, created_at) VALUES($1, $2, $3)", task.ID, task.Name, task.CreatedAt)
	return err
}

func (r *PostgresRepository) ListTasks(ctx context.Context, skip uint64, take uint64) ([]schema.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Tasks
	tasks := []schema.Task{}
	for rows.Next() {
		task := schema.Task{}
		if err = rows.Scan(&task.ID, &task.Name, &task.CreatedAt); err == nil {
			tasks = append(tasks, task)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
