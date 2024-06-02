package utils

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

const passCost = 14
const secretEnvName = "JWT_SECRET"

type CustomScanner interface {
	Scan(dest ...any) error
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Password  string `json:"-"`
}

type CreateUserRequestBody struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"required,gte=18,lte=120"`
	Password  string `json:"password" validate:"required"`
}

type GetUsersRequestBody struct {
	IDs []string `json:"ids" validate:"required,dive"`
}

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), passCost)

	return string(bytes), err
}

func IsPasswordMatching(pass, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

func InitEnv() error {
	err := godotenv.Load(".env")

	return err
}

func GetSecret() (string, bool) {
	return os.LookupEnv(secretEnvName)
}
