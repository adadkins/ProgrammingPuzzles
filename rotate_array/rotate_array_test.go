package rotate_array_test

import (
	"programmingpuzzles/rotate_array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray_in_place(t *testing.T) {
	//[1, 2, 3, 4, 5, 6] rotated by two becomes [3, 4, 5, 6, 1, 2].
	t.Run("Rotate by 2 example", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{3, 4, 5, 6, 1, 2}
		result := rotate_array.Rotate_array_in_place(input, 2)
		assert.Equal(t, expected, result)
	})

	t.Run("Rotate by length of the array so no movement", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{1, 2, 3, 4, 5, 6}
		result := rotate_array.Rotate_array_in_place(input, 6)
		assert.Equal(t, expected, result)
	})

	t.Run("Rotate by much longer array by larger rotation", func(t *testing.T) {
		input := []int{10, 15, 20, 25, 60, 70, 80, 90, 100, 110, 110, 130, 140, 150, 1, 2, 3, 4, 5, 6}
		expected := []int{110, 130, 140, 150, 1, 2, 3, 4, 5, 6, 10, 15, 20, 25, 60, 70, 80, 90, 100, 110}
		result := rotate_array.Rotate_array_in_place(input, 10)
		assert.Equal(t, expected, result)
	})
}

func TestRotateArray_new_array(t *testing.T) {
	//[1, 2, 3, 4, 5, 6] rotated by two becomes [3, 4, 5, 6, 1, 2].
	t.Run("Rotate by 2 example", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{3, 4, 5, 6, 1, 2}
		result := rotate_array.Rotate_array_in_array(input, 2)
		assert.Equal(t, expected, result)
	})

	t.Run("Rotate by length of the array so no movement", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := []int{1, 2, 3, 4, 5, 6}
		result := rotate_array.Rotate_array_in_array(input, 6)
		assert.Equal(t, expected, result)
	})

	t.Run("Rotate by much longer array by larger rotation", func(t *testing.T) {
		input := []int{10, 15, 20, 25, 60, 70, 80, 90, 100, 110, 110, 130, 140, 150, 1, 2, 3, 4, 5, 6}
		expected := []int{110, 130, 140, 150, 1, 2, 3, 4, 5, 6, 10, 15, 20, 25, 60, 70, 80, 90, 100, 110}
		result := rotate_array.Rotate_array_in_array(input, 10)
		assert.Equal(t, expected, result)
	})
}
