package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidIdsFromRange(t *testing.T) {
	testCases := []struct {
		input    [2]int
		expected []int
	}{
		{[2]int{10, 22}, []int{11, 22}},
		{[2]int{95, 115}, []int{99}},
		{[2]int{998, 1012}, []int{1010}},
		{[2]int{1188511880, 1188511890}, []int{1188511885}},
		{[2]int{222220, 222224}, []int{222222}},
		{[2]int{824824821, 824824827}, []int{}},
		{[2]int{2121212118, 2121212124}, []int{}},
		// manual cases
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v-%v", tc.input[0], tc.input[1])
		t.Run(name, func(t *testing.T) {
			result := InvalidIdsFromRange(tc.input[0], tc.input[1])
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIsValidId(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"123123", false},
		{"222222", false},
		{"222223", true},

		{"10", true},
		{"11", false},
		{"12", true},

		{"0101", true},
		{"0", true},

		{"22", false},
		{"1010", false},
		{"222222", false},
		{"446446", false},
		{"38593859", false},

		{"2121212118", true},
		{"2121212119", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsValidId(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
