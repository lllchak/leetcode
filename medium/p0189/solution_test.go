package p0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(&input1, 3)
	want1 := []int{5, 6, 7, 1, 2, 3, 4}

	input2 := []int{-1, -100, 3, 99}
	rotate(&input2, 2)
	want2 := []int{3, 99, -1, -100}

	assert.Equal(t, want1, input1)
	assert.Equal(t, want2, input2)
}
