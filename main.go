package main

import (
	"log"
	"net/http"
	"time"

    dbUtils "budget-be/db"
    "budget-be/controllers"

	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := dbUtils.ConnectToDb()

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()

	e := echo.New()
	ctrl := &controllers.BaseController{
        DB: db,
	}

	e.Use(middleware.Logger())
	e.Validator = &controllers.CustomValidator{Validator: validator.New()}

	e.GET("/user/:id", ctrl.GetUser)
	e.GET("user", ctrl.GetUsers)
	e.POST("/user", ctrl.CreateUser)

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
