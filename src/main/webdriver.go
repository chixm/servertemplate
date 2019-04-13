package main

// Test the web pages automatically

import (
	"net/http"
	_ "os"
	"service"

	"github.com/gorilla/mux"
	"github.com/sclevine/agouti"
)

var webdriver *agouti.WebDriver

// Test All pages loaded when server launches
func InitializeWebdriver() {
	webdriver = agouti.ChromeDriver()

	err := webdriver.Start()
	if err != nil {
		logger.Debug(err)
	}
	defer webdriver.Stop()

	browser, err := webdriver.NewPage()
	if err != nil {
		logger.Error("Failed to open page:%v", err)
	}

	err = browser.Navigate("http://www.dmm.com/")
	if err != nil {
		logger.Error(err)
	}
	html, _ := browser.HTML()
	logger.Debug(html)
}

func webdriveAction(w http.ResponseWriter, r *http.Request) {
	var urlParams = mux.Vars(r)
	// choose webdirve command from URL
	err := service.ExecuteWebdriver(urlParams[`command`], w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
