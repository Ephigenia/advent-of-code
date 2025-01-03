package main

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
	filename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}

	parsedInput := parseInput(rawInput)

	processInputPartOne(parsedInput)
}

func parseInput(str string) []string {
	re := regexp.MustCompile(`\s+`)
	return re.Split(str, -1)
}

func iterateStone(in string) []string {
	ln := len(in)
	if in == "0" {
		return []string{"1"}
	}

	if ln%2 == 0 {
		left := strings.TrimLeft(string(in)[0:ln/2], "0")
		right := strings.TrimLeft(string(in)[ln/2:ln], "0")
		if left == "" {
			left = "0"
		}
		if right == "" {
			right = "0"
		}
		return []string{left, right}
	}

	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	str := strconv.FormatInt(int64(val*2024), 10)
	return []string{str}
}

func iterate(in []string) []string {
	ret := []string{}
	for _, stone := range in {
		ret = append(ret, iterateStone(stone)...)
	}
	return ret
}

func processInputPartOne(input []string) {
	processInput(input, 5)
	// turns out running it 75 times is memory intensive
	// redesign to stream
	// processInput(input, 75)
}

func processInput(input []string, iterationsCount int) {
	totalCount := 0
	for i, val := range input {
		next := []string{val}
		// iterating number by numebr is the same as iterataing the whole set
		for j := 0; j < iterationsCount; j++ {
			next = iterate(next)
			fmt.Printf("Iteration %d/%d (count: %d)\n", i, j, len(next))
		}
		totalCount += len(next)
	}

	fmt.Printf("Total count: %d\n", totalCount)
}
