package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrMaxAndIndex(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	max, index := ArrMaxAndIndex(vals)
	assert.Equal(t, 5, max)
	assert.Equal(t, 4, index)
}
