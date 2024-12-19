package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
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

func operate(a, b int, operation string) int {
	if operation == "+" {
		return a + b
	}
	if operation == "*" {
		return a * b
	}
	panic("invalid operatoration")
}

func TryMe(result int, args ...int) int {
	solutions := []int{}
	for index, arg := range args {
		if index == len(args)-1 {
			continue
		}
		solution := operate(arg, args[index+1], "+")
		solutions = append(solutions, solution)
		solution = operate(arg, args[index+1], "*")
		solutions = append(solutions, solution)
	}
	spew.Dump(solutions)
	return 0
}

func processInputPartOne(inputs []InputLine) {
	for i, input := range inputs {
		if i > 1 {
			continue
		}
		correctCombinations := TryMe(input.result, input.numbers...)

		fmt.Printf("%d: %v %d\n", input.result, input.numbers, correctCombinations)
	}
	// spew.Dump(inputs)
}
