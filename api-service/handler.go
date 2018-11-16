package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/s-take/goecho-postgre-sample/db"
	"github.com/s-take/goecho-postgre-sample/schema"
	"github.com/segmentio/ksuid"
)

// listTasksHandler is ...
func listTasksHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "List Table")
	}
}

// getTaskHandler is ...
func getTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskid := c.Param("taskid")
		return c.String(http.StatusOK, "Task id is "+taskid)
	}
}

// CreateTaskHandler is ...
func createTaskHandler() echo.HandlerFunc {
	type TaskParam struct {
		Name string `json:"name"`
	}
	return func(c echo.Context) error {
		param := new(TaskParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		// create task
		createdAt := time.Now().UTC()
		id, err := ksuid.NewRandomWithTime(createdAt)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create task")
		}
		task := schema.Task{
			ID:        id.String(),
			Name:      param.Name,
			CreatedAt: createdAt,
		}
		if err := db.InsertTask(c, task); err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, "Failed to create task")
		}
		return c.JSON(http.StatusOK, "insert ok")
	}
}
