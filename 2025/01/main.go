package main

import (
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ephigenia/advent-of-code/2025/lib"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib.ExitWithError(errors.New("please provide a filename"))
	}
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(inputFilename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

type InputItem struct {
	direction string
	offset    int
}

func validateInputLine(line string) error {
	match, _ := regexp.MatchString("^[LR]\\d+$", line)
	if !match {
		return errors.New("line does not match expected format")
	}

	return nil
}

func convertInputLineToItem(line string) []InputItem {
	if err := validateInputLine(line); err != nil {
		panic(err)
	}

	value, _ := strconv.Atoi(line[1:])

	return []InputItem{
		{
			direction: line[0:1],
			offset:    value,
		},
	}
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")

	startPosition := 0

	position := startPosition

	for i, line := range lines {
		newItems := convertInputLineToItem(line)

		newPosition := calculateNewPosition(position, newItems[0].direction, newItems[0].offset)
		spew.Dump(i, position, newItems[0], newPosition)
	}
}

func calculateNewPosition(position int, direction string, offset int) int {
	return position + offset
}
