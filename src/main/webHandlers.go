package main

// author chixm

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

// URI of endpoints
const (
	URI_LOGIN          = `/login`
	URI_USER_INFO      = `/userInfo`
	URI_INFORMATION    = `/information`
	URI_MATCHING       = `/match/{roomId}`
	URI_SUBMIT_LOGIN   = `/submitLogin`
	URI_WEBDRIVER      = `/browser/{command}`
	URU_REQUEST_VIEWER = `/request/view`
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
	services[URU_REQUEST_VIEWER] = requestViewHander
	return &services, nil
}

// "HomeHandler HTMLテンプレートによるWebページの表示"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "../../resources/top.html", "../../resources/parts/header.html")
}

// 各URLごとの処理を記述
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "../../resources/information.html", "../../resources/parts/header.html")
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : going to implement websocket chat.
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "../../resources/login.html", "../../resources/parts/header.html")
}

// Write out current Request headers and params to know what is sent to server.
// This function is to develop client application or test API from external application.
func requestViewHander(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	if _, err := bufbody.ReadFrom(r.Body); err != nil {
		log.Println(err)
	}
	body := bufbody.String()
	disp := make(map[string]interface{})
	disp[`method`] = r.Method
	disp[`header`] = r.Header
	disp[`body`] = body
	showTemplate(w, disp, "../../resources/httprequest.html", "../../resources/parts/header.html")
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
	t, err := template.ParseFiles(htmlFilesInResource...)
	if err != nil {
		log.Println(err)
		http.Error(w, `Error on Parsing Template`, http.StatusInternalServerError)
		return
	}
	t.Execute(w, values)
}
