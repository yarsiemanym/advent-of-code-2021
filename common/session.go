package common

import "os"

var Session string

func InitSession() {
	session := os.Getenv("AoC_Session")
	Session = session
}
