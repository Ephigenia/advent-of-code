package main

import (
	"errors"
	"fmt"
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
	filename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		levels := lib.ArrStrToInt(strings.Split(line, " "))
		fmt.Printf("#%d levels: %v\n", i, levels)
	}
}
