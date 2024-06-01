package main

import (
	"log"
	"net/http"
	"time"

	"budget-be/controllers"
	dbUtils "budget-be/db"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	echoJwt "github.com/labstack/echo-jwt/v4"
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
    e.GET("/login", ctrl.Login)

    authorized := e.Group("/authorized")

    config := echoJwt.Config{
        NewClaimsFunc: func(c echo.Context) jwt.Claims {
            return new(controllers.JwtCustomClaims)
        },
        SigningKey: []byte("secret"),
    } 

    authorized.Use(echoJwt.WithConfig(config))

    authorized.GET("/test", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "result": "works!",
        })
    })

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
