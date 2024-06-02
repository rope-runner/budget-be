package controllers

import "database/sql"

type BaseController struct {
	DB *sql.DB
    Secret string
}
