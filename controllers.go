package main

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BaseController struct {
    db *sql.DB
}


type CustomValidator struct {
    validator *validator.Validate
}


func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (bc *BaseController) getUser(c echo.Context) error {
    return c.JSON(http.StatusOK, 1)
}



func (bc *BaseController) createUser(c echo.Context) error {
	//	var uQuery UserQuery
	//
	//	if err := echo.QueryParamsBinder(c).FailFast(true).String("name", &uQuery.Name).String("email", &uQuery.Email).BindError(); err != nil {
	//		return err
	//	}
	//
	//	if err := c.Validate(uQuery); err != nil {
	//		return err
	//	}
	//
	//	user := &User{
	//		Name:  uQuery.Name,
	//		Email: uQuery.Email,
	//		ID:    uuid.NewString(),
	//	}

	return c.JSON(http.StatusOK, 1)
}
