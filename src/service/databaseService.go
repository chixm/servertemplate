package service

// Service Package is an Implementation of functions used in other packages.

import (
	"github.com/jmoiron/sqlx"
)

type DatabaseConnection struct {
	Connections map[string]*sqlx.DB
}

func (d *DatabaseConnection) getDb(id string) *sqlx.DB {
	return d.Connections[id]
}

var dbConnHolder *DatabaseConnection
