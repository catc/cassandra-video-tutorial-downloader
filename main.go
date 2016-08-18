package main

import (
	"fmt"
	"strings"
)

type Tutorial struct {
	TutorialURL string // original tutorial url on datastax site
	Filename    string // name of file
	VideoURL    string
	Config      *config
}

func main() {
	// load config
	config := getConfig()

	rawTutorialUrls := []string{
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-6",
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-1",
	}

	for _, u := range rawTutorialUrls {
		split := strings.Split(u, "/")
		name := split[len(split)-1]
		t := Tutorial{
			TutorialURL: u,
			Filename:    name,
			Config:      config,
		}

		go processURL(&t)
	}

	var input string
	fmt.Scanln(&input)

	// parseVimeoSrc("133680477")

	/*
		TODO
		- create channel that goes through each tutorial struct and runs functions in order:
			- get datastax stuff and parse vimeo id
			- get vimeo page source and parse raw link
			- for all raw links, download video
	*/
}
