package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/davecgh/go-spew/spew"
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

// correct answer is 7229350537438%
func processInputPartOne(input string) {
	data, operators := parseInput(input)

	problems := make([]int, len(data[0]))
	for x := 0; x < len(data[0]); x++ {
		operator := operators[x]
		for y := 0; y < len(data); y++ {
			// fmt.Printf("data[%d][%d]=%d operator=%s\n", y, x, data[y][x], operator, problems[x])
			switch operator {
			case "+":
				problems[x] += data[y][x]
			case "*":
				if problems[x] == 0 {
					problems[x] = 1
				}
				problems[x] *= data[y][x]
			}
		}
	}

	spew.Dump(problems)
	fmt.Printf("part one %d", lib2024.ArrIntSum(problems))
}
