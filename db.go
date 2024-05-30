package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const dbEnvName = "DB_CONNECTION"

type CustomScanner interface {
	Scan(dest ...any) error
}

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	InternalId int
}

func getConnectionString() (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", errors.New(fmt.Sprintf("error loading .env file: %s", err.Error()))
	}

	connStr, present := os.LookupEnv(dbEnvName)

	if !present {
		return "", errors.New(fmt.Sprintf("missing env variable %s", dbEnvName))
	}

	log.Println(connStr)

	return connStr, nil
}

func connectToDb() (*sql.DB, error) {
	connectionStr, err := getConnectionString()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", connectionStr)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to connect to the db: %s", err.Error()))
	}

	if err := db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("database is not responding: %s", err.Error()))
	}

	return db, nil
}

func getUserFromDB(usrUuid string, db *sql.DB) (*User, error) {
	sqlStatement := `
    SELECT * FROM users WHERE uuid = $1;
    `

	var dbUser *User = &User{}

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

func getUsersFromDB(uuids []string, db *sql.DB) ([]*User, error) {
	var users []*User = make([]*User, 0, len(uuids))

	sqlStatement := `
    SELECT * FROM users where uuid in ($1)
    `

	rows, err := db.QueryContext(context.Background(), sqlStatement, strings.Join(uuids, ","))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dbUser *User = &User{}

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

func createUserInDB(db *sql.DB) {
	// TODO
}

func mapUserToStruct(scanner CustomScanner, user *User) error {
	err := scanner.Scan(&user.InternalId, &user.FirstName, &user.LastName, &user.Email, &user.ID)

	if err != nil {
		return err
	}

	return nil
}
