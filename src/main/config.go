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
func InitializeConfig() {
	logger.Println(`[Configuration]`)
	loadConfiguration()
}

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
}

// ALL Configuration File Contents Structure
type Configuration struct {
	Database []*DbConfig
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
