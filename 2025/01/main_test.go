package main

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
	dial := NewSafeDial(50)

	dial.Rotate("L", 68)
	assert.Equal(t, 82, dial.position)

	dial.Rotate("L", 30)
	assert.Equal(t, 52, dial.position)

	dial.Rotate("R", 48)
	assert.Equal(t, 0, dial.position)

	dial.Rotate("L", 5)
	assert.Equal(t, 95, dial.position)

	dial.Rotate("R", 60)
	assert.Equal(t, 55, dial.position)

	dial.Rotate("L", 55)
	assert.Equal(t, 0, dial.position)

	dial.Rotate("L", 1)
	assert.Equal(t, 99, dial.position)

	dial.Rotate("L", 99)
	assert.Equal(t, 0, dial.position)

	dial.Rotate("R", 14)
	assert.Equal(t, 14, dial.position)

	dial.Rotate("L", 82)
	assert.Equal(t, 32, dial.position)
}

func TestRotationR60To55(t *testing.T) {
	dial := NewSafeDial(95)
	dial.Rotate("R", 60)
	assert.Equal(t, 55, dial.position)
}

// func TestRotationR159To55(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 55, calculateNewPosition(95, "R", 159))
// }

// func TestRotationL82to32(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 32, calculateNewPosition(14, "L", 82))
// }
// func TestRotationL181to32(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 32, calculateNewPosition(14, "L", 181))
// }

// func TestLBelowZero(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 32, calculateNewPosition(14, "L", 82))
// }

// func TestRAboveMaxMultiple(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 2, calculateNewPosition(0, "R", 300))
// }

// func TestRAboveMax(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	assert.Equal(t, 55, calculateNewPosition(95, "R", 60))
// }
