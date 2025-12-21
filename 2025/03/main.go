package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/Ephigenia/advent-of-code/2024/lib"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib.ExitWithError(errors.New("please provide a filename"))
	}
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(inputFilename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		vals := lib.ArrStrToInt(strings.Split(line, ""))
		// numbers can not be sorted as this would change their order
		ProcessSet(vals)
	}
}

func ProcessSet(vals []int) int {

	return 0
}
