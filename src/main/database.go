package main

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// to use this function. you need to create database tables.
// Database table info is in sql directory.
// If you want to run this database, install MySQL and create database and tables in the sql directory.

var database map[string]*sqlx.DB

// init configuration of database. requires configurations loaded.
// To See MySQL driver settings format. Go to https://github.com/go-sql-driver/mysql/
func initializeDatabaseConnections() {
	database = make(map[string]*sqlx.DB)
	for _, dbConf := range config.Database {
		// Connecting to MySQL server.
		db := sqlx.MustOpen("mysql", dbConf.Username+`:`+dbConf.Password+"@tcp("+dbConf.Host+":"+strconv.Itoa(dbConf.Port)+")/"+dbConf.Name+"?characterEncoding=utf8")
		db.SetMaxIdleConns(dbConf.MaxIdle)
		db.SetMaxOpenConns(dbConf.MaxOpen)
		database[dbConf.ID] = db
		// Exec Query for test
		db.QueryRow("select 'test connection' from dual where 1 = $1", 1)
		logger.Println(`DB Connection Created for ` + dbConf.Name + " User[" + dbConf.Username + "] maxIdle::" + strconv.Itoa(dbConf.MaxIdle) + " maxOpen::" + strconv.Itoa(dbConf.MaxOpen))
	}
}

// remove connections if server ends.
func terminateDatabaseConnections() {
	for key, d := range database {
		logger.Println(`Closing Database :` + key)
		if err := d.Close(); err != nil {
			logger.Error(err)
		}
	}
}

func registerUser(userID, password string) error {
	db := database[`masterDB`]
	_, err := db.Exec(`insert into user(userID, password) value (?,?)`, userID, password)
	if err != nil {
		return err
	}
	logger.Info(`Inserting User Info [` + userID + `]`)
	return nil
}
