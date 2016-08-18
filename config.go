package main

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	AuthToken string
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
