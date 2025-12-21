package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	lib "github.com/Ephigenia/advent-of-code/2025/lib"
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

// correct answer is 17179
func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		vals := lib2024.ArrStrToInt(strings.Split(line, ""))
		// numbers can not be sorted as this would change their order
		processingResult := ProcessSet(vals)
		total += processingResult
		fmt.Printf("Processing line: %s, result: %d, total %d\n", line, processingResult, total)
	}
	fmt.Printf("Total result: %d\n", total)
}

func ProcessSet(vals []int) int {
	// find the max in the vals and then the
	firstMax, firstMaxIndex := lib.ArrMaxAndIndex(vals[0 : len(vals)-1])
	secondMax, _ := lib.ArrMaxAndIndex(vals[firstMaxIndex+1:])
	return firstMax*10 + secondMax
}
