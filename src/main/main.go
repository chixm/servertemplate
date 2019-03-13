package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// サーバー本体
var server *http.Server

func main() {

	initialize()

	launchServer(createServerEndPoints())
}

func initialize() {

}

// エンドポイントの設定
func createServerEndPoints() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/information", InformationHandler)
	r.HandleFunc("/match/{roomId}", MatchHandler)

	// /static/以下をファイルに直アクセス可能な部分として定義
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`../../static`))))
	return r
}

// サーバ起動処理
func launchServer(r http.Handler) error {
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:19090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return server.ListenAndServe()
}
