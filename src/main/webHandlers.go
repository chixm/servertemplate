package main

// author chixm

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

// URI of endpoints
const (
	uri_LOGIN                = `/login`
	uri_USER_INFO            = `/userInfo`
	uri_INFORMATION          = `/information`
	uri_MATCHING             = `/match/{roomId}`
	uri_SUBMIT_LOGIN         = `/submitLogin`
	uri_WEBDRIVER            = `/browser/{command}`
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
	services[uri_INFORMATION] = InformationHandler
	services[uri_MATCHING] = MatchHandler
	services[uri_LOGIN] = loginHandler
	services[uri_SUBMIT_LOGIN] = submitLoginHandler
	services[uri_WEBDRIVER] = WebdriverHandler
	// loginCheckInterceptor redirects to login page if user was not logged in.
	services[uri_USER_INFO] = loginCheckInterceptor(UserInfoHandler)
	services[uri_USER_REGIST] = userRegistrationHandler
	services[uri_SUBMIT_USER_REGIST] = submitUserRegistHandler
	services[uri_COMPLETE_USER_REGIST] = completeRegistUser
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

// first user registration
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
	randomKey := createUniqID()
	// define new user
	user := userBase{ID: address, Password: password}
	// user registration is ready for an hour
	if err := setRedisObject(randomKey, user, 60*60); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // TODO : make error page
		return
	}

	// create user registration url
	url := `http://` + server.Addr + uri_COMPLETE_USER_REGIST + `?hash=` + randomKey

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

/**
* This URL is accessed from Email.
 */
func completeRegistUser(w http.ResponseWriter, r *http.Request) {
	// get key from query
	requestKey := r.URL.Query().Get("hash")

	var b []byte
	var err error
	if b, err = getRedisObject(requestKey); err != nil {
		w.Write([]byte(`Failed to register user[` + err.Error() + `]`))
		return
	}
	if len(b) == 0 {
		w.Write([]byte(`User is not in confirm. Make sure to click Email link in an hour after sent.`))
		return
	}
	user := userBase{}
	if err = json.Unmarshal(b, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Failed to load User Data from temporary Redis`))
		return
	}

	// register user to database.
	if err := registerUser(user.ID, user.Password); err != nil {
		w.Write([]byte(`Failed to register user[` + err.Error() + `]`))
		return
	}
	logger.Info(`Registered User Info [` + user.ID + `]`)

	if _, err := w.Write([]byte(`User registration finished.`)); err != nil {
		panic(err)
	}
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
