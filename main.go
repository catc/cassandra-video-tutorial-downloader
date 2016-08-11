package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	// regex stuff
	"errors"
	"regexp"
)

const VIMEO_PREFIX = "https://player.vimeo.com/video/"

func main() {
	// load config
	config := getConfig()

	// compile regex
	vimeoID, err := regexp.Compile(`id="vimeo-([0-9]+)"`)
	if err != nil {
		log.Fatal("error compiling regex", err)
	}
	config.VimeoIdRegex = vimeoID

	/*u := "https://fpdl.vimeocdn.com/vimeo-prod-skyfire-std-us/01/1736/5/133680477/435748879.mp4?token=57abda74_0xe6a77d6a912ad9f049cc27f81ed13e56750b5a97"
	p := "/Users/catalincovic/development/git/50-go-projects/cassandra-video-tutorial-scraper/vid"
	downloadFile(p, u)*/

	u := "https://academy.datastax.com/courses/ds220-data-modeling/introduction-introduction-killrvideo"
	getDatastaxPageSrc(config, u)
}

func getDatastaxPageSrc(c *config, url string) (err error) {
	/*
		TODO
		- rename function
		- generate upstream function that takes all urls from page and gets ids for each
			- can use channels!
	*/
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("failed to create request")
		return err
	}

	cookie := strings.Join([]string{c.Session, c.AuthToken}, "; ")
	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to get response")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("request failed to return OK")
		return err
	}

	// ready body and convert to string
	data, err := ioutil.ReadAll(res.Body)
	str := string(data[:])

	regexMatch := c.VimeoIdRegex.FindStringSubmatch(str)
	if len(regexMatch) == 0 {
		return errors.New("Could not find vimeo id")
	}

	vimeoID := regexMatch[1]

	fmt.Println("vimeo id is ", vimeoID)

	return nil

	/*
		TODO - handle error in upstream function
	*/
}

func getVideoSource() {
	/*
		TODO - integrate with vimeo api
	*/
}

func downloadFile(filepath string, url string) (err error) {
	/*
		TODO
		- add support for title
		- add command line animation
			- https://github.com/tj/go-spin/blob/master/spin.go
		- add parallel downloads
	*/

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
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

/*
	TODO
	- need to find source
	- need to display ALL sources that it will get
	- need to display download bar per file?
*/

/*

	func HTTPDownload(uri string) ([]byte, error) {
    fmt.Printf("HTTPDownload From: %s.\n", uri)
    res, err := http.Get(uri)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()
    d, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ReadFile: Size of download: %d\n", len(d))
    return d, err
}

func WriteFile(dst string, d []byte) error {
    fmt.Printf("WriteFile: Size of download: %d\n", len(d))
    err := ioutil.WriteFile(dst, d, 0444)
    if err != nil {
        log.Fatal(err)
    }
    return err
}

func DownloadToFile(uri string, dst string) {
    fmt.Printf("DownloadToFile From: %s.\n", uri)
    if d, err := HTTPDownload(uri); err == nil {
        fmt.Printf("downloaded %s.\n", uri)
        if WriteFile(dst, d) == nil {
            fmt.Printf("saved %s as %s\n", uri, dst)
        }
    }
}

*/
