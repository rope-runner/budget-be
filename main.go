package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "budget"
)

func main() {
	db, err := connectToDb()

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()

	e := echo.New()
	controllers := BaseController{
		db: db,
	}

	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/user/:id", controllers.getUser)
	e.POST("/user", controllers.createUser)

	s := &http.Server{
		Addr:              ":8080",
		Handler:           e,
		ReadTimeout:       20 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      20 * time.Second,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
