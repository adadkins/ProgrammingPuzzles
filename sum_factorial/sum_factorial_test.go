package sum_factorial_test

import (
	"programmingpuzzles/sum_factorial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPrimesArray(t *testing.T) {
	t.Run("Find sums of two numbers, one prime one not prime", func(t *testing.T) {
		result := sum_factorial.SumFactorialArray([]int{3, 6})
		expected := []int{6, 9}
		assert.Equal(t, expected, result)
	})

	t.Run("Find sums of 4 numbers", func(t *testing.T) {
		result := sum_factorial.SumFactorialArray([]int{7, 10, 12})
		expected := []int{9, 27, 27}
		assert.Equal(t, expected, result)
	})

	t.Run("Find sums of numltiple powers of 10 up to 10,0000", func(t *testing.T) {
		result := sum_factorial.SumFactorialArray([]int{10, 100, 1000, 10000})
		expected := []int{27, 648, 10539, 149346}
		assert.Equal(t, expected, result)
	})
	t.Run("Find sums of numltiple powers of 10", func(t *testing.T) {
		result := sum_factorial.SumFactorialArray([]int{10})
		expected := []int{27}
		assert.Equal(t, expected, result)
	})
}
