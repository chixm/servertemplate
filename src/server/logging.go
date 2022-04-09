package server

import (
	"log"
	"os"
	"path/filepath"

	logrus "github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

var logFile *os.File

/**
 * Using Logrus Library for Logging formatter.
 * See https://github.com/sirupsen/logrus for detail.
 * This library formats log to JSON format to make it easier to read from other log analyzer.
 * useFile true:uses logfile false: outputs to standard output
 *
 */
func SetupLog(useFile bool) {
	// Configure Log Formats
	var lg = logrus.New()
	mode := int32(0777)
	os.Mkdir(`.`+string(filepath.Separator)+`log`, os.FileMode(mode))
	file, err := os.OpenFile(`./log/application.log`, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(mode))
	if err != nil {
		log.Fatal(err)
	}
	f := new(logrus.JSONFormatter)
	f.TimestampFormat = "2006-01-02T15:04:05.999Z07:00"
	lg.Formatter = f

	hostname, _ := os.Hostname()
	Logger = lg.WithField("host", hostname) //always write log with hostname.

	if useFile {
		lg.SetOutput(file)
	} else {
		lg.SetOutput(os.Stdout)
	}
	Logger.Info("Logrus is Setup for logging.")
}

func TerminateLog() {
	logFile.Close()
}
