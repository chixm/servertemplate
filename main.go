package servertemplate

import (
	"net/http"
	"strconv"
	"time"

	src "github.com/chixm/servertemplate/src/server"
	"github.com/gorilla/mux"
)

const useLogFile = false // If this param is true, write out log to /log/application.log file instead of writing to stdout.

func main() {
	initialize()
	defer terminate()

	// server launch
	launchServer(createServerEndPoints())
}

// initialize systems around this server.
func initialize() {
	src.SetupLog(useLogFile)

	src.InitializeConfig()

	src.InitializeUniqueIDMaker()

	src.InitializeDatabaseConnections()

	src.InitializeRedis()

	src.InitializeWebdriver()

	src.InitializeEmailSender()

	src.InitializeBatch()

	initializeServiceFunctions()
}

func terminate() {
	src.TerminateDatabaseConnections()

	src.TerminateBatch()

	src.TerminateRedis()

	src.TerminateLog()
}

// settings of endpoints
func createServerEndPoints() *mux.Router {
	r := mux.NewRouter()
	serviceMap, err := src.LoadServices()
	if err != nil {
		panic(`Coud not define URL of Server service.` + err.Error())
	}

	// load all handlers listed in service.webHandlers.
	for url, handler := range *serviceMap {
		r.HandleFunc(url, handler)
	}

	// load default web socket
	src.InitializeWebSocket()
	r.HandleFunc("/ws", src.Ws)

	// Files under /static can accessed by /static/(filename)...
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(`/resources/static`))))
	return r
}

// launch server
func launchServer(r http.Handler) error {
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + strconv.Itoa(src.Config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	src.Logger.Info("Execute Server ::" + server.Addr)
	return server.ListenAndServe()
}

/** give initialized functions and connections to service. */
func initializeServiceFunctions() {

}
