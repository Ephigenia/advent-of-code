package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"

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

	processInputPartTwo(rawInput)
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

func processInputPartTwo(rawInput string) {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := regex.FindAllStringSubmatch(rawInput, -1)
	spew.Dump(matches)
	sum := 0
	active := true
	for _, match := range matches {
		command := match[0]
		if command == "do()" {
			active = true
			continue
		} else if command == "don't()" {
			active = false
			continue
		} else if command[:3] == "mul" {
			if active {
				values := []string{match[1], match[2]}
				factors := lib.ArrStrToInt(values)
				sum += factors[0] * factors[1]
			}
		}
	}

	fmt.Printf("Result is %d\n", sum)
}
