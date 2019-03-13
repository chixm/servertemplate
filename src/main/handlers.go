package main

import (
	"html/template"
	"log"
	"net/http"
)

// 各URLごとの処理を記述

// HTMLテンプレートによるWebページの表示
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../resources/top.html", "../../resources/parts/header.html")
	if err != nil {
		log.Println(err)
		http.Error(w, `Error on Template`, http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func InformationHandler(w http.ResponseWriter, r *http.Request) {

}

func MatchHandler(w http.ResponseWriter, r *http.Request) {

}
