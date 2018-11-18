package db

import (
	"github.com/labstack/echo"
	"github.com/s-take/goecho-postgre-sample/schema"
)

// Repository ...
type Repository interface {
	Close()
	InsertTask(c echo.Context, task schema.Task) error
	ListTasks(c echo.Context) ([]schema.Task, error)
	GetTask(c echo.Context, id string) (schema.Task, error)
	DeleteTask(c echo.Context, id string) error
	UpdateTask(c echo.Context, id string, name string) error
}

var impl Repository

// SetRepository ...
func SetRepository(repository Repository) {
	impl = repository
}

// Close ...
func Close() {
	impl.Close()
}

// InsertTask ...
func InsertTask(c echo.Context, task schema.Task) error {
	return impl.InsertTask(c, task)
}

// ListTasks ...
func ListTasks(c echo.Context) ([]schema.Task, error) {
	return impl.ListTasks(c)
}

// GetTask ...
func GetTask(c echo.Context, id string) (schema.Task, error) {
	return impl.GetTask(c, id)
}

// DeleteTask ...
func DeleteTask(c echo.Context, id string) error {
	return impl.DeleteTask(c, id)
}

// UpdateTask ...
func UpdateTask(c echo.Context, id string, name string) error {
	return impl.UpdateTask(c, id, name)
}
