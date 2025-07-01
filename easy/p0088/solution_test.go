package p0088

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1Lhs := []int{1, 2, 3, 0, 0, 0}
	input1Rhs := []int{2, 5, 6}

	merge(&input1Lhs, 3, input1Rhs, 3)
	want1 := []int{1, 2, 2, 3, 5, 6}

	input2Lhs := []int{1}
	input2Rhs := []int{}

	merge(&input2Lhs, 1, input2Rhs, 0)
	want2 := []int{1}

	input3Lhs := []int{0}
	input3Rhs := []int{1}

	merge(&input3Lhs, 0, input3Rhs, 1)
	want3 := []int{1}

	assert.Equalf(t, want1, input1Lhs, "inputLhs = %v, inputRhs = %v", input1Lhs, input1Rhs)
	assert.Equalf(t, want2, input2Lhs, "inputLhs = %v, inputRhs = %v", input2Lhs, input2Rhs)
	assert.Equalf(t, want3, input3Lhs, "inputLhs = %v, inputRhs = %v", input3Lhs, input3Rhs)
}
