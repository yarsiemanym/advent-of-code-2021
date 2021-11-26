package common

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func ReadFile(path string) string {
	log.Debugf("Reading file \"%v\".", path)
	bytes, err := ioutil.ReadFile(path)
	Check(err)
	log.Debugf("%v bytes read from file.", len(bytes))
	log.Debug("Converting bytes to text.")
	text := string(bytes)
	log.Debugf("Text is %v characters long.", len(text))
	log.Tracef("Text is \"%v\".", Peek(text, PEEK_MAX_DEFAULT))
	return text
}
