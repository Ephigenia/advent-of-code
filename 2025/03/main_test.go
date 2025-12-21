package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessSet(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	ref := ProcessSet(vals)
	assert.Equal(t, "45", ref)
}
