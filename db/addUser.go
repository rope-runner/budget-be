package db

import (
	"context"
	"database/sql"

	"budget-be/utils"

	"github.com/google/uuid"
)

func CreateUserInDB(db *sql.DB, user *utils.CreateUserRequestBody) error {
	sqlStatement := `
    INSERT INTO users (uuid, first_name, last_name, email, age, password)
    VALUES ($1, $2, $3, $4, $5, $6);
    `

	hashed, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = db.ExecContext(context.Background(), sqlStatement, uuid.NewString(), user.FirstName, user.LastName, user.Email, user.Age, hashed)

	if err != nil {
		return err
	}

	return nil
}
