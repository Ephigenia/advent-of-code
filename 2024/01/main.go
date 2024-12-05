package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		ExitWithError(errors.New("please provide a filename"))
	}
	filename := path.Join(pwd, args[1])

	rawInput, err := ReadInputFile(filename)
	if err != nil {
		ExitWithError(err)
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
	for _, line := range lines {
		values := ArrStrToInt(strings.Split(line, "   "))
		spew.Dump(values)
	}
}
