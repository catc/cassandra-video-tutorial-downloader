package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	// ---
	"log"
	"path/filepath"
	"strings"
)

const VIMEO_PREFIX = "https://player.vimeo.com/video/"

// var VIMEO_SOURCE_REGEX = regexp.MustCompile(`\(function\(e,a\)\{var t\=(.+?)\;if`)
var VIMEO_SOURCE_REGEX = regexp.MustCompile(`\"progressive\"\:(.+?)\]`)

type VideoSources struct {
	Height int
	Url    string
}

func parseVimeoSrc(id string) ([]VideoSources, error) {
	var sources = []VideoSources{}

	url := VIMEO_PREFIX + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("failed to create request")
		return sources, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to get response")
		return sources, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("request failed to return OK")
		return sources, err
	}

	// ready body and convert to string
	data, err := ioutil.ReadAll(res.Body)
	str := string(data[:])

	// parse content into json strings
	match := VIMEO_SOURCE_REGEX.FindString(str)
	if match == "" {
		return sources, errors.New("Could not find 'progressive' object containing video sources")
	}
	// parse out `progressive: ` to get array of video sources
	match = strings.TrimLeft(match, "\"progressive:\"")

	bytes := []byte(match)

	err = json.Unmarshal(bytes, &sources)
	if err != nil {
		return sources, err
	}

	return sources, nil
}

// func downloadFile(filename string, url string) (err error) {
func downloadFile(t *Tutorial) (err error) {
	/*
		TODO
		- add support for title
		- add command line animation
			- https://github.com/tj/go-spin/blob/master/spin.go
			- https://github.com/sethgrid/multibar/blob/master/multibar.go
			- http://stackoverflow.com/questions/30532886/golang-dynamic-progressbar
			- https://github.com/cheggaaa/pb
			- https://github.com/mitchellh/ioprogress
			- http://stackoverflow.com/questions/22421375/how-to-print-the-bytes-while-the-file-is-being-downloaded-golang
		- get download size from content length header
			- https://www.socketloop.com/tutorials/golang-get-download-file-size
			- http://code.runnable.com/VJHbrd73QVQn_ifr/go-print-all-http-headers-from-a-url-for-golang-and-download
		- add parallel downloads
			- http://cavaliercoder.com/blog/downloading-large-files-in-go.html
			- https://github.com/cavaliercoder/grab
	*/

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fp := filepath.Join(dir, "vids", t.Filename+".mp4")

	// Create the file
	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(t.VideoURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
