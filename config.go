package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
)

type config struct {
	Session      string
	AuthToken    string
	VimeoIdRegex *regexp.Regexp
}

func getConfig() *config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	conf := &config{}
	if err = decoder.Decode(conf); err != nil {
		log.Fatal("error loading config.json", err)
	}
	return conf
}
