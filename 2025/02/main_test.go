package main

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func IsValidId(id int) bool {
	re := regexp.MustCompile(`^(.+)(?:\\1)+$`)

	idStr := strconv.Itoa(id)
	r := re.Match([]byte(idStr))

	spew.Dump(r)
	return r
}

// 1012

func TestValidId(t *testing.T) {
	assert.True(t, IsValidId(12))
	assert.True(t, IsValidId(13))
}

func TestInvalId(t *testing.T) {
	assert.False(t, IsValidId(11))
	assert.False(t, IsValidId(22))
	assert.False(t, IsValidId(1010))
	assert.False(t, IsValidId(222222))
	assert.False(t, IsValidId(446446))
	assert.False(t, IsValidId(38593859))
}
