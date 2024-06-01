package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"budget-be/utils"
)

func GetUserFromDB(usrUuid string, db *sql.DB) (*utils.User, error) {
	sqlStatement := `
    SELECT * FROM users WHERE uuid = $1;
    `

	var dbUser *utils.User = &utils.User{}

	row := db.QueryRowContext(context.Background(), sqlStatement, usrUuid)

	err := mapUserToStruct(row, dbUser)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New(fmt.Sprintf("no user with id: %s", usrUuid))
	case err != nil:
		return nil, err
	default:
		return dbUser, nil
	}
}
