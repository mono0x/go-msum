package msum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	repeat := 10000
	part := []float64{1, 1e100, 1, -1e100}
	input := make([]float64, 0, len(part)*repeat)
	for i := 0; i < repeat; i++ {
		input = append(input, part...)
	}
	assert.Equal(t, 20000.0, Sum(input))
}
