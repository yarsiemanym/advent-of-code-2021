package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

func ensureInputExists(day int) string {
	log.Debug("Checking existence of input file.")
	path := fmt.Sprintf("./day%02d/input.txt", day)
	log.Tracef("path = \"%v\"", path)
	_, err := os.Open(path)

	if err != nil {
		log.Debug("Local copy of input file not found.")
		url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)
		download(url, path)
	} else {
		log.Debug("Local copy of input file found.")
	}

	return path
}

func download(url string, path string) {
	log.Debug("Downloading input file.")
	log.Tracef("url = \"%v\"", url)
	log.Tracef("path = \"%v\"", path)

	if len(common.Session) == 0 {
		log.Panic("Cannot download puzzle input because the AOC_SESSION environment variable is not set.")
	}

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

	log.Debugf("%v bytes downloaded.", size)
	log.Debug("Input file saved.")
}
