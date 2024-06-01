package controllers

import (
	dbUtils "budget-be/db"
	"net/http"

	"github.com/labstack/echo/v4"
)
func (bc *BaseController) GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := dbUtils.GetUserFromDB(id, bc.DB)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
