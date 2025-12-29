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
		// {[2]int{998, 1012}, []int{1010}},
		// {[2]int{1188511880, 1188511890}, []int{1188511885}},
		// {[2]int{222220, 222224}, []int{222220}},
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
		{"0101", true},
		{"0", true},

		{"11", false},
		{"22", false},
		{"1010", false},
		{"222222", false},
		{"446446", false},
		{"38593859", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsValidId(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestStrContainsRepeatedPattern(t *testing.T) {
	testCases := []struct {
		input           string
		expectedPattern string
		expectedCount   int
	}{
		{"11", "11", 1},
		{"22", "22", 1},
		{"1010", "1010", 1},
		{"222222", "222222", 1},
		{"446446", "446446", 1},
		{"38593859", "38593859", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			pattern, count := StrContainsRepeatedPattern(tc.input)
			assert.Equal(t, tc.expectedPattern, pattern)
			assert.Equal(t, tc.expectedCount, count)
		})
	}
}
