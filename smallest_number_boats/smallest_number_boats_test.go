package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestBoatsCalculator(t *testing.T) {
	t.Run("Test given example in problem", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 200, 150, 80}, 200)
		assert.Equal(t, 3, result)
	})

	t.Run("Test example one person per boat", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 100, 100, 100, 100}, 100)
		assert.Equal(t, 5, result)
	})

	t.Run("Test given example in problem", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 200, 150, 80, 30, 30, 30, 30}, 200)
		assert.Equal(t, 4, result)
	})

	t.Run("Test example all person in one boat", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 100, 100, 100, 100}, 600)
		assert.Equal(t, 1, result)
	})

	t.Run("Test example no one fits in the boat at all", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 100, 100, 100, 100}, 10)
		assert.Equal(t, 0, result)
	})

	t.Run("Test example some heavier people cant fit in a boat at all", func(t *testing.T) {
		result := SmallestBoatCalculator([]int{100, 100, 100, 100, 100, 210, 210}, 200)
		assert.Equal(t, 3, result)
	})
}
