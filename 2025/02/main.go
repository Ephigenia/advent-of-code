package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/Ephigenia/advent-of-code/2025/lib"
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
	processInputPartTwo(rawInput)
}

// right answer is 13919717792
func processInputPartOne(input string) {
	re := regexp.MustCompile("\r?\n")
	normalizedInput := re.ReplaceAllString(input, "")
	parts := strings.Split(normalizedInput, ",")
	invalidIds := []int{}
	for _, part := range parts {
		splitted := lib2024.ArrStrToInt(strings.Split(part, "-"))
		invalidIds = append(invalidIds, InvalidIdsFromRange(splitted[0], splitted[1])...)
	}
	sum := lib2024.ArrIntSum(invalidIds)
	fmt.Printf("Part one: %d\n", sum)
}

func processInputPartTwo(input string) {
	re := regexp.MustCompile("\r?\n")
	normalizedInput := re.ReplaceAllString(input, "")
	parts := strings.Split(normalizedInput, ",")
	invalidIds := []int{}
	for _, part := range parts {
		splitted := lib2024.ArrStrToInt(strings.Split(part, "-"))
		invalidIds = append(invalidIds, InvalidIdsFromRange2(splitted[0], splitted[1])...)
	}
	sum := lib2024.ArrIntSum(invalidIds)
	fmt.Printf("Part Two: %d\n", sum)
}

func IsValidId(str string) bool {
	// ids starting with zero are invalid
	if str[0:1] == "0" {
		return true
	}
	size := len(str) / 2
	chunks := []string{
		str[0:size],
		str[size:],
	}
	if chunks[0] == chunks[1] {
		return false
	}
	return true
}

func InvalidIdsFromRange(start, end int) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= end; i++ {
		if !IsValidId(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func IsValidId2(str string) bool {
	// ids starting with zero are invalid
	if str[0:1] == "0" {
		return true
	}
	pattern, count := StrRepeatedPattern(str)
	return !(count >= 2 && strings.Repeat(pattern, count) == str)
}

func StrRepeatedPattern(str string) (pattern string, count int) {
	// Check for repeated patterns (11, 1010, 446446, etc.)
	for patternLen := 1; patternLen <= len(str)/2; patternLen++ {
		pattern := str[:patternLen]
		occurrences := len(str) / patternLen
		// fmt.Printf("str %s %d %s\n", str, patternLen, pattern)
		if strings.Repeat(pattern, occurrences) == str[:len(pattern)*(occurrences)] {
			return pattern, strings.Count(str, pattern)
		}
	}
	return "", 0
}

func InvalidIdsFromRange2(start, end int) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= end; i++ {
		if !IsValidId2(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
