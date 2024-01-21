package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractAbbrev_lax(t *testing.T) {
	line, abbrev := extractAbbrev("Los Angeles     LAX")
	assert.Equal(t, "Los Angeles", line)
	assert.Equal(t, "LAX", abbrev)
}

func Test_extractAbbrev_wi(t *testing.T) {
	line, abbrev := extractAbbrev("Wisconsin  \tWI")
	assert.Equal(t, "Wisconsin", line)
	assert.Equal(t, "WI", abbrev)
}

func Test_extractAbbrev_empty(t *testing.T) {
	line, abbrev := extractAbbrev("")
	assert.Equal(t, "", line)
	assert.Equal(t, "", abbrev)
}

func Test_reverse(t *testing.T) {
	answer := reverse("test")
	assert.Equal(t, "tset", answer)
}
