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

func backtrack(target, index, current int, numbers []int, concat bool) bool {
	if index == len(numbers) {
		return current == target
	}
	if backtrack(target, index+1, current+numbers[index], numbers, concat) {
		return true
	}
	if backtrack(target, index+1, current*numbers[index], numbers, concat) {
		return true
	}

	// part two
	if concat {
		val, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, numbers[index]))
		if backtrack(target, index+1, val, numbers, concat) {
			return true
		}
	}

	return false
}

func TryMe(result int, numbers []int, concat bool) bool {
	return backtrack(result, 1, numbers[0], numbers, concat)
}

func processInputPartTwo(inputs []InputLine) {

}

func processInputPartOne(inputs []InputLine) {
	total := 0
	total2 := 0
	for _, input := range inputs {
		if TryMe(input.result, input.numbers, false) {
			total += input.result
		}

		if TryMe(input.result, input.numbers, true) {
			total2 += input.result
		}

		fmt.Printf("%d: %v\n", input.result, input.numbers)
	}
	fmt.Printf("Total 1part: %d\n", total)
	fmt.Printf("Total 2part: %d\n", total2)
	// spew.Dump(inputs)
}
