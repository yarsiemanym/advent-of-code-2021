package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

type Puzzle struct {
	Year      int
	Day       int
	InputFile string
	solution  Solution
}

func (puzzle *Puzzle) SetSolution(solution Solution) {
	puzzle.solution = solution
}

func (puzzle Puzzle) Solve() Answer {

	log.Info("Ensuring input file exists.")
	puzzle.EnsureInputFileExists()
	log.Info("Input file exists.")

	if puzzle.solution == nil {
		log.Panicf("Solution is not set.")
	}

	if !puzzle.IsUnlocked() {
		log.Panicf("Day %v has not been unlocked.", puzzle.Day)
	}

	log.Info("Solving puzzle.")
	answer := puzzle.solution(puzzle)
	log.Info("Puzzle solved!")

	return answer
}

func (puzzle Puzzle) IsUnlocked() bool {
	log.Debugf("Checking if day %v has been unlocked.", puzzle.Day)

	est, err := time.LoadLocation(("EST"))
	Check(err)

	var puzzleUnlockAt time.Time
	if puzzle.Day != 0 {
		puzzleUnlockAt = time.Date(puzzle.Year, 11, 30, 0, 0, 0, 0, est).Add(time.Hour * 24 * time.Duration(puzzle.Day))
	}
	log.Tracef("puzzleUnlockAt = \"%v\"", puzzleUnlockAt)

	isUnlocked := puzzleUnlockAt.Before(time.Now())
	log.Tracef("isUnlocked = %v", isUnlocked)

	return isUnlocked
}

func (puzzle *Puzzle) EnsureInputFileExists() string {
	log.Debug("Checking existence of input file.")
	target := fmt.Sprintf("./day%02d/input.txt", puzzle.Day)
	log.Tracef("target = \"%v\"", target)

	_, err := os.Open(target)

	if err != nil {
		log.Debug("Local copy of input file not found.")
		url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", puzzle.Year, puzzle.Day)
		downloadInputFile(url, target)
	} else {
		log.Debug("Local copy of input file found.")
	}

	puzzle.InputFile = target
	return target
}

func downloadInputFile(source string, target string) {
	log.Debug("Downloading input file.")
	log.Tracef("source = \"%v\"", source)
	log.Tracef("target = \"%v\"", target)

	if len(Session) == 0 {
		log.Panic("Cannot download puzzle input because the AOC_SESSION_TOKEN environment variable is not set.")
	}

	file, err := os.Create(target)

	if os.IsNotExist(err) {
		dir := path.Dir(target)
		os.Mkdir(dir, 0766)
		file, err = os.Create(target)
	}

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
