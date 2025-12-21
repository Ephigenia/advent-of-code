package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessSet(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	ref := ProcessSet(vals)
	assert.Equal(t, 45, ref)
}

func TestProcessSetEdge(t *testing.T) {
	vals := []int{0, 0, 0, 0, 5}
	ref := ProcessSet(vals)
	assert.Equal(t, 5, ref)
}
func TestProcessSetEdge2(t *testing.T) {
	vals := []int{5, 0, 0, 0, 0}
	ref := ProcessSet(vals)
	assert.Equal(t, 50, ref)
}

func TestProcessSetEdge3(t *testing.T) {
	vals := []int{9, 9, 9, 0, 9}
	ref := ProcessSet(vals)
	assert.Equal(t, 99, ref)
}

func TestProcessSet2(t *testing.T) {
	vals := []int{2, 1, 3, 1, 1}
	ref := ProcessSet(vals)
	assert.Equal(t, 31, ref)
}

var tableTestData = []struct {
	in       []int
	expected int
}{
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, 98},
	{[]int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, 89},
	{[]int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, 78},
	{[]int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, 92},
}

func TestProcessSetTable(t *testing.T) {
	for _, tt := range tableTestData {
		testName := "Input" + fmt.Sprint(tt.in)
		t.Run(testName, func(t *testing.T) {
			result := ProcessSet(tt.in)
			assert.Equal(t, tt.expected, result)
		})
	}
}
