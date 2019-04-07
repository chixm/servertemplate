package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// instance of the server
var server *http.Server

const USE_LOG_FILE = false // If this param is true, write out log to /log/application.log file instead of writing to stdout.

func main() {
	initialize()

	launchServer(createServerEndPoints())

	defer terminate()
}

func initialize() {
	SetupLog(USE_LOG_FILE)

	InitializeConfig()

	InitializeUniqueIdMaker()

	InitializeDatabaseConnections()

	InitializeAutotest() //optional
}

func terminate() {
	TerminateDatabaseConnections()

	TerminateLog()
}

// URI of endpoints
const (
	URI_LOGIN        = `/login`
	URI_USER_INFO    = `/userInfo`
	URI_INFORMATION  = `/information`
	URI_MATCHING     = `/match/{roomId}`
	URI_SUBMIT_LOGIN = `/submitLogin`
)

// settings of endpoints
func createServerEndPoints() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc(URI_INFORMATION, InformationHandler)
	r.HandleFunc(URI_MATCHING, MatchHandler)
	r.HandleFunc(URI_LOGIN, LoginHandler)
	r.HandleFunc(URI_SUBMIT_LOGIN, SubmitLoginHandler)

	// loginCheckInterceptor redirects to login page if user was not logged in.
	r.HandleFunc(URI_USER_INFO, loginCheckInterceptor(UserInfoHandler))

	// Files under /static can accessed by /static/(filename)...
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`../../static`))))
	return r
}

// launch server
func launchServer(r http.Handler) error {
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:19090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Execute Server ::" + server.Addr)
	return server.ListenAndServe()
}
