package main

import (
	"fmt"
)

type Tutorial struct {
	TutorialURL  string
	VimeoID      string
	Filename     string
	VideoSources []struct {
		Height int
		Url    string
	}
}

func main() {
	// load config
	config := getConfig()

	rawTutorialUrls := []string{
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-6",
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-1",
	}

	var tutorials = []Tutorial{}
	for _, u := range rawTutorialUrls {
		tutorials = append(tutorials, Tutorial{
			TutorialURL: u,
		})
	}

	fmt.Println(config)

	// parseVimeoSrc("133680477")

	/*
		TODO
		- create channel that goes through each tutorial struct and runs functions in order:
			- get datastax stuff and parse vimeo id
			- get vimeo page source and parse raw link
			- for all raw links, download video
	*/
}
