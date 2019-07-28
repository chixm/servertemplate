package main

// author chixm

import (
	"html/template"
	"log"
	"net/http"
)

// URI of endpoints
const (
	URI_LOGIN        = `/login`
	URI_USER_INFO    = `/userInfo`
	URI_INFORMATION  = `/information`
	URI_MATCHING     = `/match/{roomId}`
	URI_SUBMIT_LOGIN = `/submitLogin`
	URI_WEBDRIVER    = `/browser/{command}`
)

/**
 * Define Webservice URL and web handler methods.
 */
func LoadServices() (*(map[string]func(w http.ResponseWriter, r *http.Request)), error) {
	services := make(map[string]func(w http.ResponseWriter, r *http.Request))
	services[`/`] = HomeHandler
	services[URI_INFORMATION] = InformationHandler
	services[URI_MATCHING] = MatchHandler
	services[URI_LOGIN] = LoginHandler
	services[URI_SUBMIT_LOGIN] = SubmitLoginHandler
	services[URI_WEBDRIVER] = WebdriverHandler
	// loginCheckInterceptor redirects to login page if user was not logged in.
	services[URI_USER_INFO] = loginCheckInterceptor(UserInfoHandler)

	return &services, nil
}

// "HomeHandler HTMLテンプレートによるWebページの表示"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/top.html", "/parts/header.html")
}

// 各URLごとの処理を記述
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/information.html", "/parts/header.html")
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : going to implement websocket chat.
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/login.html", "/parts/header.html")
}

func WebdriverHandler(w http.ResponseWriter, r *http.Request) {

}

func SubmitLoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate login

	setLoginCookie(&w, r)
	log.Println(`Login Cookie has been set`)

	// If user succeeded to login. Go to UserInfo page.
	http.Redirect(w, r, URI_USER_INFO, http.StatusFound)
}

// Must be logged In.
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(`User Reached to The page Handler with logged in state.`)
}

/** Load HTML template in resources directoroy. */
func showTemplate(w http.ResponseWriter, values interface{}, htmlFilesInResource ...string) {
	var files []string
	for _, html := range htmlFilesInResource {
		files = append(files, config.ResourcePath+html)
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, `Error on Parsing Template`, http.StatusInternalServerError)
		return
	}
	t.Execute(w, values)
}
