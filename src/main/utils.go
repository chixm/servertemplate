package main

import (
	"io/ioutil"
	"os"
)

func stringMatches(var1, var2 string) bool {
	if var1 == `` || var2 == `` {
		logger.Info(`empty string of var1:` + var1 + ` var2:` + var2)
		return false
	}
	if var1 == var2 {
		return true
	}
	return false
}

func readTextFile(fileInResource string) (string, error) {
	f, err := os.Open(config.ResourcePath + fileInResource)
	if err != nil {
		return ``, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b), err
}
