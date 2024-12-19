package main

import (
	"errors"
	"os"
	"path"
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

	processInputPartOne(rawInput)
}

func parseMatrix(input, rowDevider string) [][]int {
	if rowDevider == "" {
		rowDevider = "|"
	}
	lines := strings.Split(input, "\n")
	rows := make([][]int, len(lines))
	for i, line := range lines {
		rows[i] = lib.ArrStrToInt(strings.Split(line, rowDevider))
	}
	return rows
}

func processInputPartOne(input string) {
	parts := strings.Split(input, "\n\n")
	if len(parts) < 2 {
		panic("invalid input")
	}

	orderingRules := parseMatrix(parts[0], "|")
	updates := parseMatrix(parts[1], ",")

	spew.Dump(orderingRules, updates)
}

func CheckOrderOfPages(pages []int, orderingRules [][]int) bool {
	return true
}
