package main

import (
	"html/template"
	"log"
	"net/http"
)

// 各URLごとの処理を記述

// HTMLテンプレートによるWebページの表示
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "../../resources/top.html", "../../resources/parts/header.html")
}

func InformationHandler(w http.ResponseWriter, r *http.Request) {

}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : going to implement websocket chat.
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "../../resources/login.html", "../../resources/parts/header.html")
}

func SubmitLoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate login

	setLoginCookie(&w, r)
	log.Println(`Login Cookie has been set`)

	http.Redirect(w, r, "/userInfo", http.StatusFound)
}

// Must be logged In.
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {

}

func showTemplate(w http.ResponseWriter, values interface{}, htmlFilesInResource ...string) {
	t, err := template.ParseFiles(htmlFilesInResource...)
	if err != nil {
		log.Println(err)
		http.Error(w, `Error on Parsing Template`, http.StatusInternalServerError)
		return
	}
	t.Execute(w, values)
}
