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

	hit := 0

	initialPosition := 50
	dial := NewSafeDial(initialPosition)

	for _, line := range lines {
		instruction := convertInputLineToItem(line)
		direction := instruction[0].direction
		offset := instruction[0].offset

		// new implementation
		dial.Rotate(direction, offset)
		if dial.IsZeroPosition() {
			hit++
		}
	}

	fmt.Printf("hits: %d\n", hit)
}

type SafeDial struct {
	min             int // minimum position
	max             int // max position
	initialPosition int // initial start position
	position        int // current position
	lastPosition    int // last position before last rotation
}

func NewSafeDial(initialPosition int) *SafeDial {
	return &SafeDial{
		min:             0,
		max:             100,
		initialPosition: initialPosition,
		position:        initialPosition,
		lastPosition:    initialPosition,
	}
}

func (s *SafeDial) IsZeroPosition() bool {
	return s.position%(s.max-1) == 0
}

func NormalizeOffset(direction string, offset int) int {
	switch direction {
	case "L":
		return -offset
	case "R":
		return +offset
	default:
		panic("invalid direction")
	}
}

func (s *SafeDial) Rotate(direction string, offset int) *SafeDial {
	s.lastPosition = s.position

	s.position += NormalizeOffset(direction, offset)

	fullRotations := s.calculateTotalRotations(s.position)
	if s.position < s.min {
		s.position = int(fullRotations)*s.max + s.position
	} else if s.position > s.max {
		s.position = s.position - int(fullRotations)*s.max
	}

	if s.position == s.max {
		s.position = 0
	}
	return s
}

func (s *SafeDial) calculateTotalRotations(newPosition int) float64 {
	if newPosition < s.min {
		return -math.Floor(float64(newPosition) / float64(s.max))
	}
	return math.Floor(float64(newPosition) / float64(s.max))
}
