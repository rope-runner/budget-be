package controllers

import (
	"budget-be/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
    User *utils.User
    jwt.RegisteredClaims
}

func (bc *BaseController) Login(c echo.Context) error {
    creds := new(utils.LoginRequestBody)

    if err := (&echo.DefaultBinder{}).BindBody(c, creds); err != nil {
        return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
    }

    if err := c.Validate(creds); err != nil {
        return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
    }

    return c.JSON(http.StatusOK, 1)
}
