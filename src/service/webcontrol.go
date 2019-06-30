package service

// Tasks Using webdriver

import (
	"net/http"

	_ "github.com/sclevine/agouti"
)

const (
	LOGIN = "login"
)

func ExecuteWebdriver(command string, w http.ResponseWriter) error {
	var err error
	err = nil
	switch command {
	case LOGIN:
		err = login(w)
	}
	return err
}

func login(w http.ResponseWriter) error {
	w.Write(wl(`start login`))
	err := webdriver.Start()
	if err != nil {
		w.Write(wl(err.Error()))
	}
	return err
}

func wl(line string) []byte {
	//wl stands for write line
	return []byte(line)
}
