package common

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func EnsureInputExists(year int, day int) string {
	log.Debug("Checking existence of input file.")
	target := fmt.Sprintf("./day%02d/input.txt", day)
	log.Tracef("target = \"%v\"", target)

	_, err := os.Open(target)

	if err != nil {
		log.Debug("Local copy of input file not found.")
		url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
		Download(url, target)
	} else {
		log.Debug("Local copy of input file found.")
	}

	return target
}

func Download(source string, target string) {
	log.Debug("Downloading input file.")
	log.Tracef("source = \"%v\"", source)
	log.Tracef("target = \"%v\"", target)

	if len(Session) == 0 {
		log.Panic("Cannot download puzzle input because the AOC_SESSION_TOKEN environment variable is not set.")
	}

	file, err := os.Create(target)
	Check(err)

	client := http.Client{}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    Session,
		Domain:   ".adventofcode.com",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	req, err := http.NewRequest("GET", source, nil)
	Check(err)

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	Check(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	Check(err)

	defer file.Close()

	log.Debugf("%v bytes downloaded.", size)
	log.Debug("Input file saved.")
}
