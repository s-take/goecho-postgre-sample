package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/s-take/goecho-postgre-sample/db"
	"github.com/s-take/goecho-postgre-sample/retry"
)

// Config of Other Service
type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func main() {
	// Read Config
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
		repo, err := db.NewPostgres(addr)
		if err != nil {
			log.Println(err)
			return err
		}
		db.SetRepository(repo)
		return nil
	})
	defer db.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/tasks", listTasksHandler())
	e.POST("/tasks", createTaskHandler())
	e.GET("/tasks/:taskid", getTaskHandler())
	e.DELETE("/tasks/:taskid", deleteTaskHandler())
	e.PUT("/tasks/:taskid", updateTaskHandler())

	// Start server
	e.Start(":8080")
}
