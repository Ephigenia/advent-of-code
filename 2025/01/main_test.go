package aoc_2025_01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInputLineToItem(t *testing.T) {
	data := []struct {
		input    string
		expected []InputItem
	}{
		{"L68", []InputItem{{direction: "L", offset: 68}}},
		{"L30", []InputItem{{direction: "L", offset: 30}}},
		{"R48", []InputItem{{direction: "R", offset: 48}}},
		{"L5", []InputItem{{direction: "L", offset: 5}}},
	}
	for _, d := range data {
		assert.Equal(t, convertInputLineToItem(d.input), d.expected)
	}
}

func TestCalculateNewPosition(t *testing.T) {
	assert.Equal(t, calculateNewPosition(50, "L", 68), 82)
	assert.Equal(t, calculateNewPosition(82, "L", 30), 52)
	assert.Equal(t, calculateNewPosition(52, "R", 48), 0)
	assert.Equal(t, calculateNewPosition(0, "L", 5), 95)
	assert.Equal(t, calculateNewPosition(95, "R", 60), 55)
	assert.Equal(t, calculateNewPosition(55, "L", 55), 0)
	assert.Equal(t, calculateNewPosition(0, "L", 1), 99)
	assert.Equal(t, calculateNewPosition(99, "L", 99), 0)
	assert.Equal(t, calculateNewPosition(0, "R", 14), 14)
	assert.Equal(t, calculateNewPosition(14, "L", 82), 32)
}

func TestIt(t *testing.T) {
	assert.Equal(t, 32, calculateNewPosition(14, "L", 82))
}
