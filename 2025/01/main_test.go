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
	// number of rotations according to the example
	dial := NewSafeDial(50)

	dial.Rotate("L", 68)
	assert.Equal(t, 82, dial.position)
	assert.Equal(t, 1, dial.zeroCrossed)

	dial.Rotate("L", 30)
	assert.Equal(t, 52, dial.position)
	assert.Equal(t, 1, dial.zeroCrossed)

	dial.Rotate("R", 48)
	assert.Equal(t, 0, dial.position)
	assert.Equal(t, 2, dial.zeroCrossed)

	dial.Rotate("L", 5)
	assert.Equal(t, 95, dial.position)
	assert.Equal(t, 2, dial.zeroCrossed)

	dial.Rotate("R", 60)
	assert.Equal(t, 55, dial.position)
	assert.Equal(t, 3, dial.zeroCrossed)

	dial.Rotate("L", 55)
	assert.Equal(t, 0, dial.position)
	assert.Equal(t, 4, dial.zeroCrossed)

	dial.Rotate("L", 1)
	assert.Equal(t, 99, dial.position)
	assert.Equal(t, 4, dial.zeroCrossed)

	dial.Rotate("L", 99)
	assert.Equal(t, 0, dial.position)
	assert.Equal(t, 5, dial.zeroCrossed)

	dial.Rotate("R", 14)
	assert.Equal(t, 14, dial.position)
	assert.Equal(t, 5, dial.zeroCrossed)

	// dial.Rotate("L", 82)
	// assert.Equal(t, 32, dial.position)
	// assert.Equal(t, 6, dial.zeroCrossed)
}

// func TestRotationWithZeroOffset(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	dial.Rotate("R", 0)
// 	assert.Equal(t, 50, dial.position)
// 	assert.Equal(t, 0, dial.zeroCrossed)
// 	dial.Rotate("L", 0)
// 	assert.Equal(t, 50, dial.position)
// 	assert.Equal(t, 0, dial.zeroCrossed)
// }

// func TestSimpleRotation(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	dial.Rotate("R", 1)
// 	assert.Equal(t, 51, dial.position)
// 	assert.Equal(t, 0, dial.zeroCrossed)
// }

// func TestZeroPosition(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	dial.Rotate("L", 50)
// 	assert.Equal(t, 0, dial.position)
// 	assert.True(t, dial.IsZeroPosition())
// }

// func TestMultipleRotations(t *testing.T) {
// 	dial := NewSafeDial(50)
// 	dial.Rotate("R", 200)
// 	assert.Equal(t, 50, dial.position)
// 	assert.Equal(t, 2, dial.zeroCrossed)
// }

// func TestRotationR60To55(t *testing.T) {
// 	dial := NewSafeDial(95)
// 	dial.Rotate("R", 60)
// 	assert.Equal(t, 55, dial.position)
// 	assert.Equal(t, 1, dial.zeroCrossed)
// }

// func TestRotationR160To55(t *testing.T) {
// 	dial := NewSafeDial(95)
// 	dial.Rotate("R", 160)
// 	assert.Equal(t, 55, dial.position)
// 	assert.Equal(t, 2, dial.zeroCrossed)
// }

// func TestMe(t *testing.T) {
// 	dial := NewSafeDial(33)
// 	dial.Rotate("L", 684)
// 	assert.Equal(t, 49, dial.position)
// }

// func TestSimpleFullRightRotations(t *testing.T) {
// 	dial := NewSafeDial(0)
// 	dial.Rotate("R", 200)
// 	assert.Equal(t, 0, dial.position)
// 	assert.Equal(t, 2, dial.zeroCrossed)
// }

// func TestSimpleFullRightRotationsPlusONe(t *testing.T) {
// 	dial := NewSafeDial(0)
// 	dial.Rotate("R", 201)
// 	assert.Equal(t, 1, dial.position)
// 	assert.Equal(t, 2, dial.zeroCrossed)
// }

// func TestSimpleFullLeftRotations(t *testing.T) {
// 	dial := NewSafeDial(0)
// 	dial.Rotate("L", 200)
// 	assert.Equal(t, 0, dial.position)
// 	assert.Equal(t, 2, dial.zeroCrossed)
// }

// func TestSimpleFullLeftRotationsPlusOne(t *testing.T) {
// 	dial := NewSafeDial(0)
// 	dial.Rotate("L", 201)
// 	assert.Equal(t, 99, dial.position)
// 	assert.Equal(t, 3, dial.zeroCrossed)
// }
