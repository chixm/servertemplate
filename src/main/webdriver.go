package main

// Test the web pages automatically

import (
	_ "os"

	"github.com/sclevine/agouti"
)

var webdriver *agouti.WebDriver

// Test All pages loaded when server launches
func initializeWebdriver() {
	webdriver = agouti.ChromeDriver()

	err := webdriver.Start()
	if err != nil {
		logger.Debug(err)
	}
	defer webdriver.Stop()

	browser, err := webdriver.NewPage()
	if err != nil {
		logger.Error("Failed to open page")
		logger.Error(err)
	}

	err = browser.Navigate("https://google.com/")
	if err != nil {
		logger.Error(err)
	}
	html, _ := browser.HTML()
	logger.Debug(html)
}
