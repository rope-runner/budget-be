package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"budget-be/utils"
)

func GetUsersFromDB(uuids []string, db *sql.DB) ([]*utils.User, error) {
	var users []*utils.User = make([]*utils.User, 0, len(uuids))

	sqlStatement := `
    SELECT * FROM users where uuid in ($1)
    `

	rows, err := db.QueryContext(context.Background(), sqlStatement, strings.Join(uuids, ","))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dbUser *utils.User = &utils.User{}

		err := mapUserToStruct(rows, dbUser)

		if err != nil {
			return nil, err
		}

		users = append(users, dbUser)
	}

	if err := rows.Err(); err != nil {
		log.Println(fmt.Sprintf("%+v\n", users))

		return nil, err
	}

	return users, nil
}
