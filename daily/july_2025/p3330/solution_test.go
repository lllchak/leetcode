package p3330

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := "abbcccc"
	input2 := "abcd"
	input3 := "aaaa"

	got1 := possibleStringCount(input1)
	want1 := 5

	got2 := possibleStringCount(input2)
	want2 := 1

	got3 := possibleStringCount(input3)
	want3 := 4

	assert.Equalf(t, want1, got1, "input = %s", input1)
	assert.Equal(t, want2, got2, "input = %s", input2)
	assert.Equal(t, want3, got3, "input = %s", input3)
}
