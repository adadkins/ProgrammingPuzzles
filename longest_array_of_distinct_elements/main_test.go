package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestArray(t *testing.T) {
	t.Run("Given example", func(t *testing.T) {
		expected := 5
		result := Longest_array_of_distinct_elements([]int{5, 1, 3, 5, 2, 3, 4, 1})
		assert.Equal(t, expected, result)
	})

	t.Run("All 10 unique", func(t *testing.T) {
		expected := 10
		result := Longest_array_of_distinct_elements([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		assert.Equal(t, expected, result)
	})

	t.Run("All 10 duplicates", func(t *testing.T) {
		expected := 1
		result := Longest_array_of_distinct_elements([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
		assert.Equal(t, expected, result)
	})

	t.Run("First and alast are duplicates", func(t *testing.T) {
		expected := 9
		result := Longest_array_of_distinct_elements([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1})
		assert.Equal(t, expected, result)
	})
}
