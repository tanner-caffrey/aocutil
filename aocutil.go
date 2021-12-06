package aocutil

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

/*
	Author: Tanner Caffrey

	aocutil (Advent of Code Utility) is a package of utilities to use when completing the advent of code (2021) in go

*/

// GetInputFromDay gets the input for a given Advent of Code day.
// The session cookie is required for the user.
func GetInputFromDay(day int, session_cookie string) (string, error) {
	day_url := "https://adventofcode.com/2021/day/" + strconv.Itoa(day) + "/input"

	reader, err := downloadFile(day_url, session_cookie)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: GetInputFromDay failed to download file.")
		return "", err
	}

	rc := *reader
	contents, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: GetInputFromDay failed to read file: ", err)
		return "", err
	}

	return string(contents), nil
}

func downloadFile(url string, session_cookie string) (*io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Cookie", "session="+session_cookie)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "[X] ERROR: DownloadFile failed when getting http response: \n", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println("[!] WARNING: DownloadFile response code: ", strconv.Itoa(resp.StatusCode))
	}
	if resp.Body != nil {
		return &resp.Body, nil
	}
	return nil, errors.New("DownloadFile failed due to no response body from url")
}
