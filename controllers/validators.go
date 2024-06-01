package controllers

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func validatePassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}

