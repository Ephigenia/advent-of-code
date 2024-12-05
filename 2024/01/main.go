package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib.ExitWithError(errors.New("please provide a filename"))
	}
	filename := path.Join(pwd, args[1])

	rawInput, err := ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInput(rawInput)
}

func ReadInputFile(filename string) (string, error) {
	// read the whole file
	dat, err := os.ReadFile(filename)
	if err != nil {
		return "", nil
	}
	return strings.TrimRight(string(dat), "\n"), nil
}

func processInput(input string) {
	lines := strings.Split(input, "\n")
	right := make([]int, len(lines))
	left := make([]int, len(lines))
	for i, line := range lines {
		values := lib.ArrStrToInt(strings.Split(line, "   "))
		left[i] = values[0]
		right[i] = values[1]
	}
	spew.Dump(left, right)
}
