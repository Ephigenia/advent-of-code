package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func ReadInputFile(filename string) (string, error) {
	// read the whole file
	dat, err := os.ReadFile(filename)
	if err != nil {
		return "", nil
	}
	return strings.TrimRight(string(dat), "\n"), nil
}

func ArrMin(input []int) (int, index int) {
	min := input[0]
	index = 0
	for i, v := range input {
		if v < min {
			min = v
			index = i
		}
	}
	return min, index
}

func ArrMax(input []int) (int, index int) {
	max := input[0]
	index = 0
	for i, v := range input {
		if v > max {
			max = v
			index = i
		}
	}
	return max, index
}

// remove on item of an array
func ArrIntPopIndex(input []int, index int) []int {
	return append(input[:index], input[index+1:]...)
}

func ArrStrToInt(input []string) []int {
	var output []int
	for _, v := range input {
		if v != "" {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			output = append(output, i)
		}
	}
	return output
}
