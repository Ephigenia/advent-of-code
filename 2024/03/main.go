package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"

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

func processInputPartOne(rawInput string) {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(rawInput, -1)
	sum := 0
	for _, match := range matches {
		values := []string{match[1], match[2]}
		factors := lib.ArrStrToInt(values)
		sum += factors[0] * factors[1]
	}

	fmt.Printf("Result is %d\n", sum)
}
