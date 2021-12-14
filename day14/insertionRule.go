package day14

import (
	"regexp"

	log "github.com/sirupsen/logrus"
)

type InsertionRule struct {
	target     string
	newElement string
}

func NewInsertionRule(target string, newElement string) *InsertionRule {
	targetPattern := regexp.MustCompile(`[A-Z]{2}`)
	if !targetPattern.MatchString(target) {
		log.Fatalf("\"%s\" is not a valid insertion target.", target)
	}

	elementPattern := regexp.MustCompile(`[A-Z]`)
	if !elementPattern.MatchString(newElement) {
		log.Fatalf("\"%s\" is not a valid element.", newElement)
	}

	return &InsertionRule{
		target:     target,
		newElement: newElement,
	}
}

func (rule *InsertionRule) Target() string {
	return rule.target
}

func (rule *InsertionRule) NewElement() string {
	return rule.newElement
}
