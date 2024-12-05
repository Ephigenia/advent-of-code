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

	rawInput, err := lib.ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

func processInputPartOne(input string) {
	spew.Dump(strings.Split(input, "\n"))
}
