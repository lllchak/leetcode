package p3330

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := "abbcccc"
	got1 := possibleStringCount(input1)
	want1 := 5

	input2 := "abcd"
	got2 := possibleStringCount(input2)
	want2 := 1

	input3 := "aaaa"
	got3 := possibleStringCount(input3)
	want3 := 4

	assert.Equalf(t, want1, got1, "input = %s", input1)
	assert.Equalf(t, want2, got2, "input = %s", input2)
	assert.Equalf(t, want3, got3, "input = %s", input3)
}
