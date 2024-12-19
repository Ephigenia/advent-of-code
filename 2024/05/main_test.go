package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckOrderOfPages(t *testing.T) {
	fmt.Printf("asd")
	pages := []int{75, 47, 61, 53, 29}
	orderingRules := [][]int{
		{75, 47},
		{75, 61},
		{75, 53},
		{75, 29},
	}
	result := CheckOrderOfPages(pages, orderingRules)
	assert.Equal(t, result, true)
}
