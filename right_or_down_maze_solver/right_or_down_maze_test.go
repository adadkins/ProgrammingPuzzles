package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMazeSolver(t *testing.T) {
	t.Run("BaseCase: 3x3 maze with 1 paths", func(t *testing.T) {
		input := [][]int{{0, 0, 0},
			{0, 1, 0},
			{0, 1, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 1, result)
	})
	t.Run("BaseCase: Find 3x3 maze with 2 paths", func(t *testing.T) {
		input := [][]int{{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 2, result)
	})
	t.Run("Find 3 paths for figure 8 maze", func(t *testing.T) {
		input := [][]int{{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 3, result)
	})

	t.Run("Find paths for 5x5 maze", func(t *testing.T) {
		input := [][]int{{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 22, result)
	})

	t.Run("Find zero paths for blocked maze", func(t *testing.T) {
		input := [][]int{{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 1, 0, 0, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 0, result)
	})

	t.Run("Find zero paths for blocked maze", func(t *testing.T) {
		input := [][]int{{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0}}
		result := MazeSolutions(input, 0, 0)
		assert.Equal(t, 0, result)
	})
}
