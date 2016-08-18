package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var VIMEO_ID_REGEX = regexp.MustCompile(`id="vimeo-([0-9]+)"`)

func getVimeoId(c *config, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("failed to create request")
		return "", err
	}

	cookie := c.AuthToken
	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("failed to get response")
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("request failed to return OK")
		return "", err
	}

	// ready body and convert to string
	data, err := ioutil.ReadAll(res.Body)
	str := string(data[:])

	regexMatch := VIMEO_ID_REGEX.FindStringSubmatch(str)
	if len(regexMatch) == 0 {
		return "", errors.New("Could not find vimeo id")
	}

	vimeoID := regexMatch[1]

	return vimeoID, nil
}
