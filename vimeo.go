package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"./progress"
)

const VIMEO_PREFIX = "https://player.vimeo.com/video/"

var VIMEO_SOURCE_REGEX = regexp.MustCompile(`\"progressive\"\:(.+?\])`)

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

	match := VIMEO_SOURCE_REGEX.FindStringSubmatch(str)
	if len(match) < 1 {
		return sources, errors.New("Could not find 'progressive' object containing video sources")
	}

	bytes := []byte(match[1])

	err = json.Unmarshal(bytes, &sources)
	if err != nil {
		return sources, err
	}

	return sources, nil
}

func downloadFile(t *Tutorial) (err error) {
	// generate filepath
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fp := filepath.Join(dir, "vids", t.Filename+".mp4")

	// create the file
	dist, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer dist.Close()

	// get the data
	resp, err := http.Get(t.VideoURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// get size
	downloadSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	// made custom ReadCloser
	src := &ReaderStatus{
		ReadCloser:   resp.Body, // this is the reader
		Total:        int64(downloadSize),
		Downloaded:   0,
		ProgressBars: t.ProgressBars,
		Bar:          t.Bar,
	}

	_, err = io.Copy(dist, src)
	if err != nil {
		return err
	}

	return nil
}

type ReaderStatus struct {
	io.ReadCloser
	Total      int64   // total file length
	Downloaded int64   // current number of bytes downloaded
	CurrentBar float32 // current '=' in progress bar

	ProgressBars *progress.ProgressBars
	Bar          *progress.Bar
}

func (rs *ReaderStatus) Read(p []byte) (int, error) {
	n, err := rs.ReadCloser.Read(p)

	pb := rs.ProgressBars
	bar := rs.Bar

	if err == nil {
		rs.Downloaded += int64(n)

		// only update progress based on bar length increments
		progress := float64(rs.Downloaded) / float64(rs.Total)

		totalBars := float64(bar.BarLength)
		barIncrement := 100 / totalBars

		currentIncrement := progress * totalBars // float64
		currentWithIncrement := float64(rs.CurrentBar) + barIncrement

		if currentIncrement >= currentWithIncrement {
			// update current bar status on reader
			rs.CurrentBar = float32(currentIncrement)

			// update progress bar
			pb.SetBarProps(bar, float32(progress), "")
		}
	}

	if err == io.EOF {
		// update bar status to 100%
		pb.SetBarProps(bar, 1, "Completed")

		// mark bar as done
		defer pb.CompletedBar()
	}

	return n, err
}
