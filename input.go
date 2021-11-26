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
	target := fmt.Sprintf("./day%02d/input.txt", day)
	log.Tracef("target = \"%v\"", target)

	_, err := os.Open(target)

	if err != nil {
		log.Debug("Local copy of input file not found.")
		url := fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day)
		download(url, target)
	} else {
		log.Debug("Local copy of input file found.")
	}

	return target
}

func download(source string, target string) {
	log.Debug("Downloading input file.")
	log.Tracef("source = \"%v\"", source)
	log.Tracef("target = \"%v\"", target)

	if len(common.Session) == 0 {
		log.Panic("Cannot download puzzle input because the AOC_SESSION environment variable is not set.")
	}

	file, err := os.Create(target)
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

	req, err := http.NewRequest("GET", source, nil)
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
