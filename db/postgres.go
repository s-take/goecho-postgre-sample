package db

import (
	"database/sql"

	"github.com/labstack/echo"
	// PostgreSQL driver
	_ "github.com/lib/pq"
	"github.com/s-take/goecho-postgre-sample/schema"
)

// PostgresRepository ...
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgres ...
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

// Close PostgresRepository
func (r *PostgresRepository) Close() {
	r.db.Close()
}

// InsertTask PostgresRepository
func (r *PostgresRepository) InsertTask(c echo.Context, task schema.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks(id, name, created_at) VALUES($1, $2, $3)", task.ID, task.Name, task.CreatedAt)
	return err
}

// ListTasks PostgresRepository
func (r *PostgresRepository) ListTasks(c echo.Context) ([]schema.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks ORDER BY created_at ")
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

// GetTask PostgresRepository
func (r *PostgresRepository) GetTask(c echo.Context, id string) (schema.Task, error) {
	task := schema.Task{}
	err := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.CreatedAt)
	if err != nil {
		return task, err
	}
	return task, nil
}

// DeleteTask PostgresRepository
func (r *PostgresRepository) DeleteTask(c echo.Context, id string) error {
	_, err := r.db.Exec("DELETE FROM tasks where id=$1", id)
	return err
}

// UpdateTask PostgresRepository
func (r *PostgresRepository) UpdateTask(c echo.Context, id string, name string) error {
	_, err := r.db.Exec("UPDATE tasks SET name = $2 WHERE id = $1", id, name)
	return err
}
