// A mock-puzzle to exercise all of the common bits

package day00

import (
	"fmt"
	"time"

	"github.com/yarsiemanym/advent-of-code-2021/common"
)

const shortDateFormat = "2006-01-02"

func Solve(path string) string {
	common.InitLogging()

	results := common.ParseFile(path, "\n", parseUser)

	var users []User
	for _, result := range results {
		users = append(users, result.(User))
	}

	// Do a thing

	var answer string

	for index, user := range users {
		answer = answer + fmt.Sprintf("User %v = { name: %v, email: %v, birthday: %v }\n", index, user.Name, user.Email, user.Birthday.Format(shortDateFormat))
	}

	return answer
}

func parseUser(text string) (error, interface{}) {
	if text == "" {
		return nil, nil
	}

	tokens := common.Split(text, ",")
	name := tokens[0]
	email := tokens[1]
	birthday, err := time.Parse(shortDateFormat, tokens[2])

	if err != nil {
		return err, nil
	}

	result := User{
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}

	return nil, result
}
