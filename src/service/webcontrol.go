package service

// Tasks Using webdriver

import (
	"net/http"

	"github.com/sclevine/agouti"
)

const (
	LOGIN = "login"
)

func ExecuteWebdriver(command string, w http.ResponseWriter) error {
	var err error
	err = nil
	webdriver := agouti.ChromeDriver()
	switch command {
	case LOGIN:
		err = login(webdriver, w)
	}
	return err
}

func login(driver *agouti.WebDriver, w http.ResponseWriter) error {
	w.Write(wl(`start login`))
	err := driver.Start()
	if err != nil {
		w.Write(wl(err.Error()))
	}

	return err
}

func wl(line string) []byte {
	//wl stands for write line
	return []byte(line)
}
