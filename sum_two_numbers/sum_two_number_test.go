package sum_two_numbers_test

import (
	"programmingpuzzles/sum_two_numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumTwoNumbersNaiveImplementation(t *testing.T) {
	t.Run("Given example", func(t *testing.T) {
		input := []int{10, 15, 3, 7}
		target := 17
		result := sum_two_numbers.Sum_two_numbers(input, target)
		expected := true
		assert.Equal(t, expected, result)
	})
	t.Run("Returns false if no solution", func(t *testing.T) {
		input := []int{10, 15, 3, 8}
		target := 17
		result := sum_two_numbers.Sum_two_numbers(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})
	t.Run("Larger example", func(t *testing.T) {
		input := []int{10, 15, 3, 7, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1}
		target := 1101
		result := sum_two_numbers.Sum_two_numbers(input, target)
		expected := true
		assert.Equal(t, expected, result)
	})
	t.Run("Handles empty input", func(t *testing.T) {
		input := []int{}
		target := 100
		result := sum_two_numbers.Sum_two_numbers(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})
	t.Run("Handles nil input", func(t *testing.T) {
		var input []int
		target := 100
		result := sum_two_numbers.Sum_two_numbers(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})

}

func TestSumTwoNumbersSinglePassImplementation(t *testing.T) {
	t.Run("Given example", func(t *testing.T) {
		input := []int{10, 15, 3, 7}
		target := 17
		result := sum_two_numbers.Sum_two_numbers_single_pass(input, target)
		expected := true
		assert.Equal(t, expected, result)
	})
	t.Run("Returns false if no solution", func(t *testing.T) {
		input := []int{10, 15, 3, 8}
		target := 17
		result := sum_two_numbers.Sum_two_numbers_single_pass(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})
	t.Run("Larger example", func(t *testing.T) {
		input := []int{10, 15, 3, 7, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1}
		target := 1101
		result := sum_two_numbers.Sum_two_numbers_single_pass(input, target)
		expected := true
		assert.Equal(t, expected, result)
	})
	t.Run("Handles empty input", func(t *testing.T) {
		input := []int{}
		target := 100
		result := sum_two_numbers.Sum_two_numbers_single_pass(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})
	t.Run("Handles nil input", func(t *testing.T) {
		var input []int
		target := 100
		result := sum_two_numbers.Sum_two_numbers_single_pass(input, target)
		expected := false
		assert.Equal(t, expected, result)
	})
}
