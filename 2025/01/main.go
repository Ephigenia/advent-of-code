package aoc_2025_01

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
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
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(inputFilename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
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

	startPosition := 50

	position := startPosition
	hit := 0

	// fmt.Printf("L86 %d\n", calculateNewPosition(startPosition, "L", 68))
	// fmt.Printf("L30 %d\n", calculateNewPosition(82, "L", 30))
	// fmt.Printf("L30 %d\n", calculateNewPosition(52, "R", 48))

	for i, line := range lines {
		newItems := convertInputLineToItem(line)

		position = calculateNewPosition(position, newItems[0].direction, newItems[0].offset)
		if position == 0 {
			hit++
		}
		fmt.Printf("#%d %s\t%d   %d\n", i, line, position, hit)
	}

	fmt.Printf("hits: %d\n", hit)
}

// TODO support offsets above 99
func calculateNewPosition(position int, direction string, offset int) int {
	var newPosition int
	switch direction {
	case "L":
		newPosition = position - offset
	case "R":
		newPosition = position + offset
	default:
		panic("invalid direction")
	}

	if newPosition < 0 {
		newPosition = 99 + newPosition + 1
	} else if newPosition > 99 {
		newPosition = newPosition - 99 - 1
	}
	return newPosition
}
