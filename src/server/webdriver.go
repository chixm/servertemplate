package server

// Test the web pages automatically

import (
	_ "os"

	"github.com/sclevine/agouti"
)

var webdriver *agouti.WebDriver

// Test All pages loaded when server launches
func InitializeWebdriver() {
	webdriver = agouti.ChromeDriver()

	err := webdriver.Start()
	if err != nil {
		Logger.Debug(err)
	}
	defer webdriver.Stop()

	browser, err := webdriver.NewPage()
	if err != nil {
		Logger.Error("Failed to open page")
		Logger.Error(err)
	}

	err = browser.Navigate("https://google.com/")
	if err != nil {
		Logger.Error(err)
	}
	html, _ := browser.HTML()
	Logger.Debug(html)
}
