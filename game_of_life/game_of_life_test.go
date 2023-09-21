package game_of_life_test

import (
	"programmingpuzzles/game_of_life"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameOfLife(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
		expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}
		result := *(game_of_life.Game_of_life(&input))
		assert.Equal(t, expected, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		input := [][]int{{1, 1}, {1, 0}}
		expected := [][]int{{1, 1}, {1, 1}}
		result := *(game_of_life.Game_of_life(&input))
		assert.Equal(t, expected, result)
	})
}
func TestGameOfLifeSteadyState(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
		result := game_of_life.Does_Board_Reach_Steady_State(&input)
		assert.Equal(t, true, result)
	})

	t.Run("Example 2", func(t *testing.T) {
		input := [][]int{{1, 1}, {1, 0}}
		result := game_of_life.Does_Board_Reach_Steady_State(&input)
		assert.Equal(t, true, result)
	})
}
