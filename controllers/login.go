package controllers

import (
	dbUtils "budget-be/db"
	"budget-be/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
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

	user, err := dbUtils.GetUserFromDBByEmail(creds.Email, bc.DB)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if !utils.IsPasswordMatching(creds.Password, user.Password) {
		return echo.ErrUnauthorized
	}

	t, err := GenerateToken(user, bc.Secret)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GenerateToken(user *utils.User, secret string) (string, error) {
	claims := &JwtCustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
