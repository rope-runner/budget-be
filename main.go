package main

import (
	"log"
	"net/http"
	"time"

	"budget-be/controllers"
	dbUtils "budget-be/db"
	"budget-be/utils"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := utils.InitEnv()

	secret, present := utils.GetSecret()

	if !present {
		secret = "default_secret"
	}

	db, err := dbUtils.ConnectToDb()

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()

	e := echo.New()
	ctrl := &controllers.BaseController{
		DB:     db,
		Secret: secret,
	}

	e.Use(middleware.Logger())
	e.Validator = &controllers.CustomValidator{Validator: validator.New()}

	e.GET("/user/:id", ctrl.GetUser)
	e.GET("user", ctrl.GetUsers)
	e.POST("/register", ctrl.CreateUser)
	e.GET("/login", ctrl.Login)

	authorized := e.Group("/api/v1")

	config := echoJwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(controllers.JwtCustomClaims)
		},
		SigningKey: []byte(secret),
	}

	authorized.Use(echoJwt.WithConfig(config))

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
