package service

import (
	"net/http"

	logrus "github.com/Sirupsen/logrus"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sclevine/agouti"
)

// Getting functions , Connections and Configurations from Main package
// And Passing Service Functions to Main Package.

// function to set Cookie to response
var setLoginCookie func(w *http.ResponseWriter, r *http.Request)

var isLoginCookieValid func(r *http.Request) bool

var loginCheckInterceptor func(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)

var logger *logrus.Entry

var redisConnHolder map[string]*redis.Pool

var webdriver *agouti.WebDriver

var sendEmail func(to []string, from string, message []byte) error

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

func LoadRedisConnections(r map[string]*redis.Pool) {
	redisConnHolder = r
}

func LoadWebDriver(w *agouti.WebDriver) {
	webdriver = w
}

func LoadSendEmailFunction(sm func([]string, string, []byte) error) {
	sendEmail = sm
}
