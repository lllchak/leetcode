package p0026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := []int{1, 1, 2}
	got1 := removeDuplicates(&input1)
	want1 := 2

	input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	got2 := removeDuplicates(&input2)
	want2 := 5

	input3 := []int{1, 2}
	got3 := removeDuplicates(&input3)
	want3 := 2

	assert.Equalf(t, want1, got1, "input = %v", input1)
	assert.Equalf(t, want2, got2, "input = %v", input2)
	assert.Equalf(t, want3, got3, "input = %v", input3)
}
