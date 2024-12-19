package main

import (
	"errors"
	"fmt"
	"os"
	"path"
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

	processInputPartOne(rawInput)
}

func parseMatrix(input, rowDevider string) [][]int {
	if rowDevider == "" {
		rowDevider = "|"
	}
	lines := strings.Split(input, "\n")
	rows := make([][]int, len(lines))
	for i, line := range lines {
		rows[i] = lib.ArrStrToInt(strings.Split(line, rowDevider))
	}
	return rows
}

func processInputPartOne(input string) {
	parts := strings.Split(input, "\n\n")
	if len(parts) < 2 {
		panic("invalid input")
	}

	orderingRules := parseMatrix(parts[0], "|")
	updates := parseMatrix(parts[1], ",")

	// spew.Dump(orderingRules, updates)
	// transform ordering rules so that before and after sets are easier to check
	// 0 pages that are before the page
	// 1 pages that are after the page
	rules := make([][][]int, 100)
	for _, rule := range orderingRules {
		page := rule[0]
		after := rule[1]
		if rules[page] == nil {
			rules[page] = make([][]int, 2, 100)
		}
		if rules[after] == nil {
			rules[after] = make([][]int, 2, 100)
		}

		rules[page][1] = append(rules[page][1], after)
		rules[after][0] = append(rules[after][0], page)
	}

	// found := isValid(75, 53, rules)
	// spew.Dump(found)
	validUpdates := make([][]int, 0)
	for _, update := range updates {
		result := true
		for i := 0; i < len(update)-1; i++ {
			cur := update[i]
			nxt := update[i+1]
			result = result && isValid(cur, nxt, rules)
		}
		if result {
			validUpdates = append(validUpdates, update)
		}
		fmt.Printf("update: %v: %v\n", update, result)
	}

	fmt.Println("Valid Updates")
	numbers := []int{}
	for _, validUpdate := range validUpdates {
		fmt.Printf("update: %v\n", validUpdate)
		numbers = append(numbers, GetMiddleNumber(validUpdate))
	}

	fmt.Printf("Middle Numbers: %v\n", numbers)
	fmt.Printf("Sum: %v\n", lib.ArrIntSum(numbers))
}

func GetMiddleNumber(numbers []int) int {
	len := len(numbers)
	middle := len / 2
	return numbers[middle]
}

func isValid(page, nextPage int, rules [][][]int) bool {
	if !lib.ArrIntFind(rules[nextPage][0], page) {
		return false
	}
	if !lib.ArrIntFind(rules[page][1], nextPage) {
		return false
	}
	return true
}

func IsCorrectOrder(page int, orderingRules [][]int) bool {

	return false
}

func CheckOrderOfPages(pages []int, orderingRules [][]int) bool {
	allOk := true

	return allOk
}
