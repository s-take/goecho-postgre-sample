package db

import (
	"context"

	"github.com/s-take/goecho-postgre-sample/schema"
)

type Repository interface {
	Close()
	InsertTask(ctx context.Context, task schema.Task) error
	ListTasks(ctx context.Context, skip uint64, take uint64) ([]schema.Task, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertTask(ctx context.Context, task schema.Task) error {
	return impl.InsertTask(ctx, task)
}

func ListTasks(ctx context.Context, skip uint64, take uint64) ([]schema.Task, error) {
	return impl.ListTasks(ctx, skip, take)
}
