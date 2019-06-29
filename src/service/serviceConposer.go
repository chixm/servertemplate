package service

import (
	"net/http"

	logrus "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

// Getting functions , Connections and Configurations from Main package
// And Passing Service Functions to Main Package.

// function to set Cookie to response
var setLoginCookie func(w *http.ResponseWriter, r *http.Request)

var isLoginCookieValid func(r *http.Request) bool

var loginCheckInterceptor func(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)

var logger *logrus.Entry

func LoadCookieFunctions(
	loginCookieSetFunc func(w *http.ResponseWriter, r *http.Request),
	loginCheckFunc func(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request),
	loginCookieValid func(r *http.Request) bool) {

	setLoginCookie = loginCookieSetFunc
	loginCheckInterceptor = loginCheckFunc
	isLoginCookieValid = loginCookieValid
}

func LoadDbConnections(dbConns map[string]*sqlx.DB) {
	dbConnHolder = &DatabaseConnection{Connections: dbConns}
}

func LoadLogger(l *logrus.Entry) {
	logger = l
}
