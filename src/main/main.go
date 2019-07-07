package main

import (
	"net/http"
	"service"
	"strconv"
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
	setupLog(USE_LOG_FILE)

	initializeConfig()

	initializeUniqueIdMaker()

	initializeDatabaseConnections()

	initializeRedis()

	initializeWebdriver()

	initializeEmailSender()

	initializeServiceFunctions()
}

func terminate() {
	terminateDatabaseConnections()

	terminateRedis()

	terminateLog()
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

	// load default web socket
	initializeWebSocket()
	r.HandleFunc("/ws", ws)

	// Files under /static can accessed by /static/(filename)...
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`../../static`))))
	return r
}

// launch server
func launchServer(r http.Handler) error {
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + strconv.Itoa(config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Execute Server ::" + server.Addr)
	return server.ListenAndServe()
}

/** give initialized functions and connections to service. */
func initializeServiceFunctions() {
	service.LoadLogger(logger)
	service.LoadCookieFunctions(setLoginCookie, loginCheckInterceptor, validateLoginCookie)
	service.LoadDbConnections(database)
	service.LoadRedisConnections(redisConnections)
	service.LoadSendEmailFunction(SendEmail)
}
