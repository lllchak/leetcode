package p0169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := []int{3, 2, 3}
	got1 := majorityElement(input1)
	got1Map := majorityElementMap(input1)
	want1 := 3

	input2 := []int{2, 2, 1, 1, 1, 2, 2}
	got2 := majorityElement(input2)
	got2Map := majorityElementMap(input2)
	want2 := 2

	assert.Equalf(t, want1, got1, "input = %v", input1)
	assert.Equalf(t, want2, got2, "input = %v", input2)

	assert.Equalf(t, want1, got1Map, "input = %v", input1)
	assert.Equalf(t, want2, got2Map, "input = %v", input2)
}
