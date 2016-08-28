package main

import (
	"errors"
)

// const VIDEO_RESOLUTION = 360
const VIDEO_RESOLUTION = 720

func (t *Tutorial) processURL() {
	bar := t.Bar
	pb := t.ProgressBars

	pb.SetBarProps(bar, 0, "Retrieving vimeo id")

	// get vimeo id
	vimeoid, err := getVimeoId(t.Config, t.TutorialURL)
	if err != nil {
		t.handleErr(err)
		return
	}

	// get vimeo source
	sources, err := parseVimeoSrc(vimeoid)
	if err != nil {
		t.handleErr(err)
		return
	}

	pb.SetBarProps(bar, 0, "Getting vimeo sources")

	var videoURL string
	for _, source := range sources {
		if source.Height == VIDEO_RESOLUTION {
			videoURL = source.Url
			break
		}
	}
	if videoURL == "" {
		err = errors.New("Error finding vimeo source specified resolution")
		t.handleErr(err)
		return
	}
	t.VideoURL = videoURL

	pb.SetBarProps(bar, 0, "Downloading file...")

	// download video
	err = downloadFile(t)
	if err != nil {
		t.handleErr(err)
	}
}

func (t *Tutorial) handleErr(err error) {
	// display error
	errStr := "ERROR: " + err.Error()
	t.ProgressBars.SetBarProps(t.Bar, 0, errStr)

	// mark bar as done
	defer t.ProgressBars.CompletedBar()
}
