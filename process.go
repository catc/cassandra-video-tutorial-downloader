package main

import (
	"fmt"
)

// const VIDEO_RESOLUTION = 720
const VIDEO_RESOLUTION = 360

func processURL(t *Tutorial) {
	vimeoid, err := getVimeoId(t.Config, t.TutorialURL)
	if err != nil {
		fmt.Println(err)
		/*
			TODO
			- handle error accordingly... post in terminal?
		*/
		return
	}
	fmt.Println("Vimeo id is ", vimeoid)

	sources, err := parseVimeoSrc(vimeoid)
	if err != nil {
		fmt.Println(sources)
		return
	}

	var videoURL string
	for _, source := range sources {
		if source.Height == VIDEO_RESOLUTION {
			videoURL = source.Url
			break
		}
	}
	if videoURL == "" {
		fmt.Println("Error finding vimeo source specified resolution")
		return
	}
	t.VideoURL = videoURL

	fmt.Println("DOWNLOADING FILE")
	err = downloadFile(t)
	if err != nil {
		fmt.Println(err)
	}
}
