package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
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

	parsedInput := parseInput(rawInput)

	processInputPartOne(parsedInput)
}

type InputLine struct {
	result  int
	numbers []int
}

func parseInput(str string) []InputLine {
	inputs := []InputLine{}
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		split := strings.Split(line, ":")
		result, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		input := InputLine{
			result:  result,
			numbers: lib.ArrStrToInt(strings.Split(split[1], " ")),
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func processInputPartOne(inputs []InputLine) {
	for _, input := range inputs {
		fmt.Printf("%d: %v\n", input.result, input.numbers)

	}
	// spew.Dump(inputs)
}
