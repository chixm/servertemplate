package main

import (
	"log"
	"net/http"
	"service"
)

// Simple login cookie function and cookie validation.
// replace functions in service package if you want to share cookie by Redis. This cookie only works for single server.

const SESSION = `_sessionId`

// ログイン時にセッション用クッキーを追加する
func setLoginCookie(w *http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     SESSION,
		Value:    createUniqId(),
		HttpOnly: true,
	}
	http.SetCookie(*w, &cookie)
}

func validateLoginCookie(r *http.Request) bool {
	c, err := r.Cookie(SESSION)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(`Cookie Value has ` + c.Value)
	return true
}

// Go to Login Page if not Logged In
func loginCheckInterceptor(exec func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateLoginCookie(r) {
			http.Redirect(w, r, service.URI_LOGIN, http.StatusFound) // 302 redirection
			return
		}
		exec(w, r)
	}
}
