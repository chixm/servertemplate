package main

import (
	"net/http"
	"service"
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

// initialize systems around this server.
func initialize() {
	SetupLog(USE_LOG_FILE)

	InitializeConfig()

	InitializeUniqueIdMaker()

	InitializeDatabaseConnections()

	initializeRedis()

	InitializeCommonFunctions()

	InitializeWebdriver() //optional
}

func terminate() {
	terminateDatabaseConnections()

	terminateRedis()

	TerminateLog()
}

// settings of endpoints
func createServerEndPoints() *mux.Router {
	r := mux.NewRouter()
	serviceMap, err := service.LoadServices()
	if err != nil {
		panic(`Coud not define URL of Server service.` + err.Error())
	}

	// load all handlers listed in service.webHandlers.
	for url, handler := range *serviceMap {
		r.HandleFunc(url, handler)
	}

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

/** send common functions to service. */
func InitializeCommonFunctions() {
	service.LoadCookieFunctions(setLoginCookie, loginCheckInterceptor, validateLoginCookie)
}
