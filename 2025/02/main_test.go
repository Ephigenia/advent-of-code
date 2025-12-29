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
		{"0101", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := StrContainsRepeatedPattern(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestStrContainsRepeatedPattern(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"11", true},
		{"22", true},
		{"1010", true},
		{"222222", true},
		{"446446", true},
		{"38593859", true},
		// manual cases
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := StrContainsRepeatedPattern(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
