package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib2024.ExitWithError(errors.New("please provide a filename"))
	}
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib2024.ReadInputFile(inputFilename)
	if err != nil {
		lib2024.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

func parseInput(input string) ([][]int, []string) {
	re := regexp.MustCompile(`\s+`)
	lines := strings.Split(input, "\n")

	operators := make([]string, 0)
	data := make([][]int, len(lines)-1)

	for y, line := range lines {
		columns := re.Split(line, -1)
		if y == len(lines)-1 {
			operators = columns
			break
		}
		data[y] = lib2024.ArrStrToInt(columns)
	}

	return data, operators
}

func evaluateOperationPartOne(input int, operator string, value int) int {
	switch operator {
	case "+":
		return input + value
	case "*":
		if input == 0 {
			return value
		}
		return input * value
	default:
		panic("unknown operator " + operator)
	}
}

// correct answer is 7229350537438%
func processInputPartOne(input string) {
	data, operators := parseInput(input)

	problems := make([]int, len(data[0]))
	for x := 0; x < len(data[0]); x++ {
		operator := operators[x]
		for y := 0; y < len(data); y++ {
			problems[x] = evaluateOperationPartOne(problems[x], operator, data[y][x])
		}
	}

	fmt.Printf("part one %d", lib2024.ArrIntSum(problems))
}
