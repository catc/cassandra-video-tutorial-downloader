package main

import (
	"strings"

	"./progress"
)

type Tutorial struct {
	TutorialURL string // original tutorial url on datastax site
	Filename    string // name of file
	VideoURL    string
	Config      *config

	ProgressBars *progress.ProgressBars
	Bar          *progress.Bar
}

func main() {
	// load config
	config := getConfig()

	// init progress bars
	pb := progress.New()

	rawTutorialUrls := []string{
		// add any urls here...
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-6",
		"https://academy.datastax.com/courses/ds220-data-modeling/quick-wins-challenge-1",
	}

	for _, u := range rawTutorialUrls {
		split := strings.Split(u, "/")
		name := split[len(split)-1]
		t := &Tutorial{
			TutorialURL: u,
			Filename:    name,
			Config:      config,

			// progress bar stuff
			ProgressBars: pb,
			Bar:          pb.CreateBar(name),
		}

		go t.processURL()
	}

	pb.Wait()
}
