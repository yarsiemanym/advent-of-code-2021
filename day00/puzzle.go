// A mock-puzzle to exercise all of the common bits

package day00

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

const shortDateFormat = "2006-01-02"

func Solve(path string) (string, string) {
	common.InitLogging()

	results := common.ParseFile(path, "\n", parseUser)

	var users []User
	for _, result := range results {
		users = append(users, result.(User))
	}

	var oldest User

	for index, user := range users {
		user.Age = time.Now().Sub(user.Birthday).Hours() / 24 / 365

		log.Debugf("User %v = { name: %v, email: %v, birthday: %v, age: %v }\n",
			index, user.Name, user.Email, user.Birthday.Format(shortDateFormat), user.Age)

		if user.Age > oldest.Age {
			oldest = user
		}
	}

	return oldest.Name, strconv.Itoa(int(oldest.Age))
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
