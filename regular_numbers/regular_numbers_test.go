package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularNumbers(t *testing.T) {
	t.Run("BaseCase: Find first regular number", func(t *testing.T) {
		input := 1
		result := RegularNumbers(input)
		assert.Equal(t, []int{1}, result)
		assert.Equal(t, input, len(result))
	})
	t.Run("BaseCase: Find two regular number", func(t *testing.T) {
		input := 2
		result := RegularNumbers(input)
		assert.Equal(t, []int{1, 2}, result)
		assert.Equal(t, input, len(result))
	})
	t.Run("BaseCase: Find 5 regular number", func(t *testing.T) {
		input := 5
		result := RegularNumbers(input)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
		assert.Equal(t, input, len(result))
	})

	t.Run("Find regular numbers up to 60", func(t *testing.T) {
		input := 26
		result := RegularNumbers(input)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, 40, 45, 48, 50, 54, 60}, result)
		assert.Equal(t, input, len(result))
	})

	t.Run("Find regular numbers up to 360", func(t *testing.T) {
		input := 58
		result := RegularNumbers(input)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, 40, 45, 48, 50, 54, 60, 64, 72, 75, 80, 81, 90, 96, 100, 108, 120, 125, 128, 135, 144, 150, 160, 162, 180, 192, 200, 216, 225, 240, 243, 250, 256, 270, 288, 300, 320, 324, 360}, result)
		assert.Equal(t, input, len(result))
	})
}
