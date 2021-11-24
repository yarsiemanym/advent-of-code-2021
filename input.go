package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func ensureInput(day int) string {
	path := fmt.Sprintf("./day%02d/input.txt", day)
	log.Debug("Checking existence of input file")
	log.Tracef("path = '%v'", path)
	_, err := os.Open(path)

	if err != nil {
		log.Debug("Input file does not exist")
		url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)
		download(url, path)
	} else {
		log.Debug("Using cached input file")
	}

	return path
}

func download(url string, path string) {
	log.Info("Downloading input file")
	log.Tracef("url = '%v'", url)
	log.Tracef("path = '%v'", path)

	file, err := os.Create(path)
	common.Check(err)

	client := http.Client{}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    common.Session,
		Domain:   ".adventofcode.com",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	req, err := http.NewRequest("GET", url, nil)
	common.Check(err)

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	common.Check(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	common.Check(err)

	defer file.Close()

	log.Debugf("%v bytes downloaded", size)
	log.Debug("Input file saved")
}
