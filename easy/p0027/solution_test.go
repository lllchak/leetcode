package p0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := []int{3, 2, 2, 3}
	got1 := removeElement(&input1, 3)
	want1 := 2

	input2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	got2 := removeElement(&input2, 2)
	want2 := 5

	input3 := []int{1}
	got3 := removeElement(&input3, 1)
	want3 := 0

	assert.Equalf(t, want1, got1, "input = %v", input1)
	assert.Equalf(t, want2, got2, "input = %v", input2)
	assert.Equalf(t, want3, got3, "input = %v", input3)
}
