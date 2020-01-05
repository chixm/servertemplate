package main

// author chixm

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// URI of endpoints
const (
	uri_LOGIN                = `/login`
	uri_USER_INFO            = `/userInfo`
	uri_SERVICE_WORKER       = `/sw`
	uri_WEB_WORKER           = `/ww`
	uri_INFORMATION          = `/information`
	uri_MATCHING             = `/match/{roomId}`
	uri_SUBMIT_LOGIN         = `/submitLogin`
	uri_WEBDRIVER            = `/browser/{command}`
	uri_USER_REGIST          = `/userregist`
	uri_SUBMIT_USER_REGIST   = `/submitUserRegist`
	uri_COMPLETE_USER_REGIST = `/doneUserRegist`
	uri_VIDEO_CHAT           = `/videochat`
)

/**
 * Define Webservice URL and web handler methods.
 */
func LoadServices() (*(map[string]func(w http.ResponseWriter, r *http.Request)), error) {
	services := make(map[string]func(w http.ResponseWriter, r *http.Request))
	services[`/`] = HomeHandler
	services[uri_INFORMATION] = InformationHandler
	services[uri_MATCHING] = MatchHandler
	services[uri_LOGIN] = loginHandler
	services[uri_SUBMIT_LOGIN] = submitLoginHandler
	services[uri_WEBDRIVER] = WebdriverHandler
	// loginCheckInterceptor redirects to login page if user was not logged in.
	services[uri_USER_INFO] = loginCheckInterceptor(userInfoHandler)
	services[uri_SERVICE_WORKER] = serviceWorkerHandler
	services[uri_WEB_WORKER] = webWorkerHandler
	services[uri_USER_REGIST] = userRegistrationHandler
	services[uri_SUBMIT_USER_REGIST] = submitUserRegistHandler
	services[uri_COMPLETE_USER_REGIST] = completeRegistUser
	services[uri_VIDEO_CHAT] = videoChat
	return &services, nil
}

// "HomeHandler HTMLテンプレートによるWebページの表示"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/top.html", "/parts/header.html", "/parts/ww.html", "/parts/footer.html")
}

// 各URLごとの処理を記述
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/information.html", "/parts/header.html", "/parts/footer.html")
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : going to implement websocket chat.
}

func videoChat(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/video.html", "/parts/header.html", "/parts/footer.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/login.html", "/parts/header.html", "/parts/footer.html")
}

func userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	showTemplate(w, nil, "/registration.html", "/parts/header.html", "/parts/footer.html")
}

// first user registration
func submitUserRegistHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		showErrorPage(w, err, `Unable to get your email address or password information. Try again later.`)
		return
	}
	// send confirmation mail to email address.
	address := r.FormValue("mailAddress")
	password := r.FormValue("password") // check with confPassword
	confPass := r.FormValue("confPassword")

	// TODO: check if adress is already registered from user table.

	logger.Info(`Registration Request from ` + address + " [" + password + "]")

	if !stringMatches(password, confPass) {
		showErrorPage(w, errors.New(`Password Confirmed is not correct.`), `Password must be same as confirmed ones.`)
		return
	}

	// send mail template with registration link
	var mailText []byte
	if mailText, err = readMailTemplate(`registerTemplate.txt`); err != nil {
		showErrorPage(w, err, `Error on creating Registration mail template. Try again later.`)
		return
	}
	randomKey := createUniqID()
	// define new user
	user := userBase{ID: address, Password: password}
	// user registration is ready for an hour
	if err := setRedisObject(randomKey, user, 60*60); err != nil {
		showErrorPage(w, err, `Failed to register to Data keyvalue store. Try again later.`)
		return
	}

	// create user registration url
	url := `http://` + server.Addr + uri_COMPLETE_USER_REGIST + `?hash=` + randomKey

	sendingText := strings.ReplaceAll(string(mailText), `#url#`, url)

	if err = sendEmail(address, `noreply@chixm.com`, []byte(sendingText)); err != nil {
		showErrorPage(w, err, `Failed to send registration email to address `+address)
		return
	}

	logger.Info(`Sent Registration Email to ` + address)

	var resVal = make(map[string]string)
	resVal[`address`] = address
	showTemplate(w, resVal, "/confirmRegistration.html", "/parts/header.html", "/parts/footer.html")
}

/**
* This URL is accessed from Email.
 */
func completeRegistUser(w http.ResponseWriter, r *http.Request) {
	// get key from query
	requestKey := r.URL.Query().Get("hash")

	var b []byte
	var err error
	if b, err = getRedisObject(requestKey); err != nil {
		showErrorPage(w, err, `Failed to register user`)
		return
	}

	if len(b) == 0 {
		showErrorPage(w, errors.New(`No record found.`), `User is not in confirm. Make sure to click Email link in an hour after sent.`)
		return
	}
	user := userBase{}
	if err = json.Unmarshal(b, &user); err != nil {
		showErrorPage(w, err, `Failed to load User Data from email hash`)
		return
	}

	// register user to database.
	if err := registerUser(user.ID, user.Password); err != nil {
		showErrorPage(w, err, `Failed to register user`)
		return
	}
	logger.Info(`Registered User Info [` + user.ID + `]`)

	showTemplate(w, user.ID, "finishRegistration.html", "/parts/header.html", "/parts/footer.html")
}

func WebdriverHandler(w http.ResponseWriter, r *http.Request) {

}

func submitLoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate login

	setLoginCookie(&w, r)
	log.Println(`Login Cookie has been set`)

	// If user succeeded to login. Go to UserInfo page.
	http.Redirect(w, r, uri_USER_INFO, http.StatusFound)
}

// Must be logged In.
func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(`User Reached to The page Handler with logged in state.`)
	// register service worker when user reached to logged in page.

	showTemplate(w, r, "/userInfo.html", "/parts/sw.html", "/parts/header.html", "/parts/footer.html")
}

// Service Woker File Reader for Avoid mine type error.
// https://stackoverflow.com/questions/47385171/serviceworker-the-script-has-an-unsupported-mime-type-chrome-extension
//
func serviceWorkerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/x-javascript")

	js, err := readTextFile(`/worker/chixmsw.js`)
	if err != nil {
		showErrorPage(w, err, `Failed to load Service worker file.`)
	}

	if _, err := w.Write([]byte(js)); err != nil {
		log.Println(err)
	}
}

func webWorkerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/x-javascript")

	js, err := readTextFile(`/worker/chixmww.js`)
	if err != nil {
		showErrorPage(w, err, `Failed to load Web Worker file.`)
	}

	if _, err := w.Write([]byte(js)); err != nil {
		log.Println(err)
	}
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
	err = t.Execute(w, values)
	if err != nil {
		logger.Error(err)
	}
}

func showErrorPage(w http.ResponseWriter, err error, msg string) {
	var values = make(map[string]interface{})
	values[`errMsg`] = error.Error
	values[`message`] = msg
	logger.Errorln(err)
	w.WriteHeader(http.StatusInternalServerError)
	t, e := template.ParseFiles("/error.html", "/parts/header.html", "/parts/footer.html")
	if e != nil {
		logger.Errorln(e)
	}
	if e := t.Execute(w, values); e != nil {
		logger.Errorln(e)
	}
}
