package main

import (
	"errors"
	"fmt"
	"math"
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

	for i, line := range lines {
		instruction := convertInputLineToItem(line)

		position = calculateNewPosition(position, instruction[0].direction, instruction[0].offset)
		if position%max == 0 {
			hit++
		}
		fmt.Printf("#%d\t%s%d\t%d -> %d\t\t%t\t%d\n",
			i,
			instruction[0].direction,
			instruction[0].offset,
			startPosition,
			position,
			position == 0,
			hit,
		)
		startPosition = position
	}

	fmt.Printf("hits: %d\n", hit)
}

const min = 0
const max = 100

func calculateTotalRotations(newPosition int) float64 {
	if newPosition < min {
		return -math.Floor(float64(newPosition) / float64(max))
	}
	return math.Floor(float64(newPosition) / float64(max))
}

// tried 46
// tried 45
// 1026 --- don't know

// TODO support offsets above 99
func calculateNewPosition(
	startPosition int,
	direction string,
	offset int,
) int {
	var newPosition int
	switch direction {
	case "L":
		newPosition = startPosition - offset
	case "R":
		newPosition = startPosition + offset
	default:
		panic("invalid direction")
	}

	fullRotations := calculateTotalRotations(newPosition)
	if newPosition < min {
		newPosition = int(fullRotations)*max + newPosition
	} else if newPosition > max {
		newPosition = newPosition - int(fullRotations)*max
	}

	return newPosition
}
