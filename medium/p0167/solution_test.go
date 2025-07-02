package p0167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := []int{2, 7, 11, 15}
	got1 := twoSum(input1, 9)
	want1 := []int{1, 2}

	input2 := []int{2, 3, 4}
	got2 := twoSum(input2, 6)
	want2 := []int{1, 3}

	input3 := []int{-1, 0}
	got3 := twoSum(input3, -1)
	want3 := []int{1, 2}

	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)
	assert.Equal(t, want3, got3)
}
