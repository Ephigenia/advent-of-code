package aoc_2025_01

import (
	"fmt"
	"os"
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
		return "", err
	}
	return strings.TrimRight(string(dat), "\n"), nil
}
