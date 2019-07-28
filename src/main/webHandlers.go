package main

// author chixm

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// URI of endpoints
const (
	URI_LOGIN                = `/login`
	URI_USER_INFO            = `/userInfo`
	URI_INFORMATION          = `/information`
	URI_MATCHING             = `/match/{roomId}`
	URI_SUBMIT_LOGIN         = `/submitLogin`
	URI_WEBDRIVER            = `/browser/{command}`
	uri_USER_REGIST          = `/userregist`
	uri_SUBMIT_USER_REGIST   = `/submitUserRegist`
	uri_COMPLETE_USER_REGIST = `/doneUserRegist`
)

/**
 * Define Webservice URL and web handler methods.
 */
func LoadServices() (*(map[string]func(w http.ResponseWriter, r *http.Request)), error) {
	services := make(map[string]func(w http.ResponseWriter, r *http.Request))
	services[`/`] = HomeHandler
	services[URI_INFORMATION] = InformationHandler
	services[URI_MATCHING] = MatchHandler
	services[URI_LOGIN] = loginHandler
	services[URI_SUBMIT_LOGIN] = submitLoginHandler
	services[URI_WEBDRIVER] = WebdriverHandler
	// loginCheckInterceptor redirects to login page if user was not logged in.
	services[URI_USER_INFO] = loginCheckInterceptor(UserInfoHandler)
	services[uri_USER_REGIST] = userRegistrationHandler
	services[uri_SUBMIT_USER_REGIST] = submitUserRegistHandler
	return &services, nil
}

// "HomeHandler HTMLテンプレートによるWebページの表示"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/top.html", "/parts/header.html", "/parts/footer.html")
}

// 各URLごとの処理を記述
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/information.html", "/parts/header.html", "/parts/footer.html")
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : going to implement websocket chat.
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/login.html", "/parts/header.html", "/parts/footer.html")
}

func userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/registration.html", "/parts/header.html", "/parts/footer.html")
}

func submitUserRegistHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// send confirmation mail to email address.
	address := r.FormValue("mailAddress")
	password := r.FormValue("password") // check with confPassword
	confPass := r.FormValue("confPassword")

	logger.Info(`Registration Request from ` + address + " [" + password + "]")

	if !stringMatches(password, confPass) {
		w.WriteHeader(http.StatusBadRequest) // TODO : make error page
		return
	}

	// send mail template with registration link
	var mailText []byte
	if mailText, err = readMailTemplate(`registerTemplate.txt`); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // TODO : make error page
		return
	}
	// create user registration url
	url := `http://` + server.Addr + uri_COMPLETE_USER_REGIST + `?hash=` + createUniqID()

	sendingText := strings.ReplaceAll(string(mailText), `#url#`, url)

	if err = sendEmail(address, `noreply@chixm.com`, []byte(sendingText)); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // TODO : make error page
		return
	}

	logger.Info(`Sent Registration Email to ` + address)

	var resVal = make(map[string]string)
	resVal[`address`] = address
	showTemplate(w, resVal, "/confirmRegistration.html", "/parts/header.html", "/parts/footer.html")
}

func WebdriverHandler(w http.ResponseWriter, r *http.Request) {

}

func submitLoginHandler(w http.ResponseWriter, r *http.Request) {
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
