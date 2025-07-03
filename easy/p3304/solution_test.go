package p3304

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input1 := 5
	got1 := kthCharacter(input1)
	want1 := "b"

	input2 := 10
	got2 := kthCharacter(input2)
	want2 := "c"

	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)
}
