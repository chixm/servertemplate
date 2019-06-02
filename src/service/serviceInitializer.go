package service

import (
	"net/http"
)

// Getting Main package functions
// Common functions in main packages.

// function to set Cookie to response
var setLoginCookie func(w *http.ResponseWriter, r *http.Request)

var isLoginCookieValid func(r *http.Request) bool

var loginCheckInterceptor func(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)

func LoadCookieFunctions(
	loginCookieSetFunc func(w *http.ResponseWriter, r *http.Request),
	loginCheckFunc func(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request),
	loginCookieValid func(r *http.Request) bool) {

	setLoginCookie = loginCookieSetFunc
	loginCheckInterceptor = loginCheckFunc
	isLoginCookieValid = loginCookieValid
}
