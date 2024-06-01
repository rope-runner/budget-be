package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
    dbUtils "budget-be/db"
    "budget-be/utils"
)

func (bc *BaseController) GetUsers(c echo.Context) error {
	rb := new(utils.GetUsersRequestBody)

	if err := (&echo.DefaultBinder{}).BindBody(c, rb); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(rb.IDs) == 0 {
		return c.JSON(http.StatusOK, make([]*utils.User, 0, 0))
	}

	users, err := dbUtils.GetUsersFromDB(rb.IDs, bc.DB)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "please try again later")
	}

	return c.JSON(http.StatusOK, users)
}
