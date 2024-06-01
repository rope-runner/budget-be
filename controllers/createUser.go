package controllers

import (
	"net/http"

    "budget-be/utils"
    dbUtils "budget-be/db"

	"github.com/labstack/echo/v4"
)

func (bc *BaseController) CreateUser(c echo.Context) error {
	u := new(utils.CreateUserRequestBody)

	if err := (&echo.DefaultBinder{}).BindBody(c, u); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if valid := validatePassword(u.Password); !valid {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "password is not valid")
	}

	if err := dbUtils.CreateUserInDB(bc.DB, u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error, please try again later")
	}

	return c.JSON(http.StatusOK, "user created")
}
