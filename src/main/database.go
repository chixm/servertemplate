package main

import (
	"github.com/jmoiron/sqlx"
)

// to use this function. you need to create database tables.
// database connections

type DataBase struct {
	Db *sqlx.DB
}
