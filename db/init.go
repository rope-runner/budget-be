package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const dbEnvName = "DB_CONNECTION"

func GetConnectionString() (string, error) {
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

func ConnectToDb() (*sql.DB, error) {
	connectionStr, err := GetConnectionString()

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
