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
	operators := []string{"+", "*"}
	solutions := make([]int, (len(args)-1)*len(operators))

	for i := range solutions {
		solutions[i] = args[0]
	}

	for argI, arg := range args {
		if argI == 0 {
			continue
		}
		for oI, operator := range operators {
			// i := (argI + (oI * 2)) - 1
			i := argI + oI
			fmt.Printf("---%d %d\n", argI, oI)
			fmt.Printf("#%d: %d %s %d\n", i, solutions[i], operator, arg)
			solutions[i] = operate(solutions[i], arg, operator)
		}
	}

	spew.Dump(solutions)

	return len(solutions)
}

func processInputPartOne(inputs []InputLine) {
	for i, input := range inputs {
		if i != 1 {
			continue
		}
		possibleEquationResults := TryMe(input.result, input.numbers...)

		fmt.Printf("%d: %v %d\n", input.result, input.numbers, possibleEquationResults)
	}
	// spew.Dump(inputs)
}
