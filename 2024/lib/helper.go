package lib

import (
	"fmt"
	"math"
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

func ArrIntSum(input []int) int {
	return ArrIntReduce(input, func(v int, i int, acc int) int {
		return acc + v
	})
}

func ArrIntReduce(input []int, reducer func(v int, i int, acc int) int) int {
	acc := 0
	for i, v := range input {
		acc = reducer(v, i, acc)
	}
	return acc
}

func ArrIntMap(input []int, mapper func(v int, i int) int) []int {
	acc := make([]int, len(input))
	for i, v := range input {
		acc[i] = mapper(v, i)
	}
	return acc
}

// run Math.asb on all values in an array
func ArrIntAbs(input []int) []int {
	return ArrIntMap(input, func(v int, i int) int {
		return int(math.Abs(float64(v)))
	})
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

func ArrIntCountOccurences(input []int, needle int) int {
	count := 0
	for _, v := range input {
		if v == needle {
			count++
		}
	}
	return count
}
