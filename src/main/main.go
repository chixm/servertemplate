package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const useLogFile = false // If this param is true, write out log to /log/application.log file instead of writing to stdout.

func main() {
	initialize()

	launchServer(createServerEndPoints())

	defer terminate()
}

// initialize systems around this server.
func initialize() {
	setupLog(useLogFile)

	initializeConfig()

	initializeUniqueIDMaker()

	initializeDatabaseConnections()

	initializeRedis()

	initializeWebdriver()

	initializeEmailSender()

	initializeBatch()

	initializeServiceFunctions()
}

func terminate() {
	terminateDatabaseConnections()

	terminateBatch()

	terminateRedis()

	terminateLog()
}

// settings of endpoints
func createServerEndPoints() *mux.Router {
	r := mux.NewRouter()
	serviceMap, err := LoadServices()
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
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`../../resources/static`))))
	return r
}

// launch server
func launchServer(r http.Handler) {
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + strconv.Itoa(config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Execute Server ::" + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/** give initialized functions and connections to service. */
func initializeServiceFunctions() {

}
