package main

// load Configuration File in root path.
// config.json for default
//

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Config Parameter Holder
var config *Configuration

// read configuration file.
func initializeConfig() {
	logger.Println(`[Configuration]`)
	loadConfiguration()
}

// loading configuration file
// configuration file should be placed in root with name 'config.json' or 'config.production.json'
func loadConfiguration() {
	// first argument of binary is environment parameter.
	if len(os.Args) > 1 {
		environment := os.Args[1]
		viper.SetConfigName(`config.` + environment)
	} else {
		viper.SetConfigName(`config`)
	}

	viper.AddConfigPath(`../../`)
	viper.AddConfigPath(`.`)

	c := Configuration{}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	printConfiguration(&c)

	config = &c
}

func printConfiguration(c *Configuration) {
	for _, v := range c.Database {
		logger.Println(`Loaded Database Configuration of ::` + v.Id + "[" + v.Host + ":" + strconv.Itoa(v.Port) + "]")
	}
	for _, r := range c.Redis {
		logger.Println(`Loaded Redis Configuration of ::` + r.Id + "[" + r.Host + "]")
	}

}

// ALL Configuration File Contents Structure
type Configuration struct {
	Port     int            // server port
	Database []*DbConfig    //database configuration
	Redis    []*RedisConfig //redis configuration
	Email    *EmailConfig   // mail configuration
}

// basic database configuration
type DbConfig struct {
	Id       string
	Name     string
	Host     string
	Port     int
	Username string
	Password string
	MaxIdle  int
	MaxOpen  int
}

type RedisConfig struct {
	Id        string // redis connection identifier
	Host      string //redis server (ip or domain)
	Port      int    // redis port
	MaxIdle   int    // connection Idle max count
	MaxActive int    // connections Active limit
}

type EmailConfig struct {
	Smtp         string // smtp server
	SmtpSvr      string // smtp server to access
	User         string // userid for authorization
	Password     string // password for authorization
	TestSendAddr string // sends email to this address for test when server launched.
}
