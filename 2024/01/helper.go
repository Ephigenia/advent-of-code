package main

import (
	"fmt"
	"os"
	"strconv"
)

func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
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
