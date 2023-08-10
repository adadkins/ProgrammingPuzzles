package sudoku_solver_test

import (
	"programmingpuzzles/sudoku_solver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	t.Run("BFS: Simple Sudoku that can be solved with no brute forcing", func(t *testing.T) {
		data := [][]int{
			{0, 0, 0, 2, 6, 0, 7, 0, 1},
			{6, 8, 0, 0, 7, 0, 0, 9, 0},
			{1, 9, 0, 0, 0, 4, 5, 0, 0},
			{8, 2, 0, 1, 0, 0, 0, 4, 0},
			{0, 0, 4, 6, 0, 2, 9, 0, 0},
			{0, 5, 0, 0, 0, 3, 0, 2, 8},
			{0, 0, 9, 3, 0, 0, 0, 7, 4},
			{0, 4, 0, 0, 5, 0, 0, 3, 6},
			{7, 0, 3, 0, 1, 8, 0, 0, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)

		expected := [][]int{
			{4, 3, 5, 2, 6, 9, 7, 8, 1},
			{6, 8, 2, 5, 7, 1, 4, 9, 3},
			{1, 9, 7, 8, 3, 4, 5, 6, 2},
			{8, 2, 6, 1, 9, 5, 3, 4, 7},
			{3, 7, 4, 6, 8, 2, 9, 1, 5},
			{9, 5, 1, 7, 4, 3, 6, 2, 8},
			{5, 1, 9, 3, 2, 6, 8, 7, 4},
			{2, 4, 8, 9, 5, 7, 1, 3, 6},
			{7, 6, 3, 4, 1, 8, 2, 5, 9},
		}
		assert.Equal(t, expected, result)
	})
	t.Run("BFS: Medium Sudoku", func(t *testing.T) {
		data := [][]int{
			{0, 2, 6, 0, 3, 0, 0, 0, 8},
			{9, 0, 0, 6, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 9, 0, 4, 0},
			{0, 0, 7, 3, 0, 2, 0, 0, 0},
			{0, 0, 4, 0, 7, 0, 8, 0, 0},
			{0, 0, 0, 8, 0, 6, 7, 0, 0},
			{0, 5, 0, 7, 2, 0, 0, 0, 0},
			{0, 0, 9, 0, 0, 5, 0, 0, 4},
			{4, 0, 0, 0, 6, 0, 2, 1, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)
		expected := [][]int{
			{1, 2, 6, 4, 3, 7, 5, 9, 8},
			{9, 4, 3, 6, 5, 8, 1, 2, 7},
			{7, 8, 5, 2, 1, 9, 3, 4, 6},
			{8, 6, 7, 3, 9, 2, 4, 5, 1},
			{3, 9, 4, 5, 7, 1, 8, 6, 2},
			{5, 1, 2, 8, 4, 6, 7, 3, 9},
			{6, 5, 1, 7, 2, 4, 9, 8, 3},
			{2, 3, 9, 1, 8, 5, 6, 7, 4},
			{4, 7, 8, 9, 6, 3, 2, 1, 5},
		}
		assert.Equal(t, expected, result)
	})
	t.Run("BFS: Hard Sudoku", func(t *testing.T) {
		data := [][]int{
			{0, 0, 6, 3, 0, 7, 0, 0, 0},
			{0, 0, 4, 0, 0, 0, 0, 0, 5},
			{1, 0, 0, 0, 0, 6, 0, 8, 2},
			{2, 0, 5, 0, 3, 0, 1, 0, 6},
			{0, 0, 0, 2, 0, 0, 3, 0, 0},
			{9, 0, 0, 0, 7, 0, 0, 0, 4},
			{0, 5, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 8, 1, 0, 9, 0, 4, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)
		expected := [][]int{
			{5, 8, 6, 3, 2, 7, 4, 9, 1},
			{7, 2, 4, 8, 9, 1, 6, 3, 5},
			{1, 9, 3, 5, 4, 6, 7, 8, 2},
			{2, 4, 5, 9, 3, 8, 1, 7, 6},
			{8, 6, 7, 2, 1, 4, 3, 5, 9},
			{9, 3, 1, 6, 7, 5, 8, 2, 4},
			{4, 5, 2, 7, 6, 3, 9, 1, 8},
			{3, 1, 9, 4, 8, 2, 5, 6, 7},
			{6, 7, 8, 1, 5, 9, 2, 4, 3},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("BFS: Handles an empty Sudoku starting state", func(t *testing.T) {
		t.Skip()
		data := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)

		expected := [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 7, 8, 9, 1, 2, 3},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{2, 3, 1, 6, 7, 4, 8, 9, 5},
			{8, 7, 5, 9, 1, 2, 3, 6, 4},
			{6, 9, 4, 5, 3, 8, 2, 1, 7},
			{3, 1, 7, 2, 6, 5, 9, 4, 8},
			{5, 4, 2, 8, 9, 7, 6, 3, 1},
			{9, 6, 8, 3, 4, 1, 5, 7, 2},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("BFS: Handles an impossible sudoku with no solution", func(t *testing.T) {
		data := [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 2},
			{0, 1, 0, 0, 0, 0, 0, 2, 0},
			{0, 0, 1, 0, 0, 0, 2, 0, 0},
			{0, 0, 0, 1, 0, 2, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 1, 0, 0, 0},
			{0, 0, 2, 0, 0, 0, 1, 0, 0},
			{0, 2, 0, 0, 0, 0, 0, 1, 0},
			{2, 0, 0, 0, 0, 0, 0, 0, 1},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)

		expected := [][]int(nil)

		assert.Equal(t, expected, result)
	})

	t.Run("BFS: Ian G's sudoku", func(t *testing.T) {
		data := [][]int{
			{1, 0, 0, 4, 3, 0, 0, 0, 0},
			{0, 2, 0, 0, 0, 6, 7, 0, 0},
			{5, 3, 0, 0, 0, 0, 0, 2, 0},
			{4, 6, 0, 0, 0, 7, 5, 0, 0},
			{0, 0, 0, 6, 0, 5, 0, 0, 0},
			{0, 0, 7, 2, 0, 0, 0, 6, 4},
			{0, 4, 0, 0, 0, 0, 0, 5, 7},
			{0, 0, 5, 1, 0, 0, 0, 3, 0},
			{0, 0, 0, 0, 7, 9, 0, 0, 6},
		}
		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Breadth_first_search(&puzzle)
		expected := [][]int{
			{1, 7, 9, 4, 3, 2, 6, 8, 5},
			{8, 2, 4, 9, 5, 6, 7, 1, 3},
			{5, 3, 6, 7, 1, 8, 4, 2, 9},
			{4, 6, 2, 3, 8, 7, 5, 9, 1},
			{9, 1, 8, 6, 4, 5, 3, 7, 2},
			{3, 5, 7, 2, 9, 1, 8, 6, 4},
			{6, 4, 1, 8, 2, 3, 9, 5, 7},
			{7, 9, 5, 1, 6, 4, 2, 3, 8},
			{2, 8, 3, 5, 7, 9, 1, 4, 6},
		}
		assert.Equal(t, expected, result)
	})
}

func TestDFS(t *testing.T) {
	t.Run("DFS: Simple Sudoku that can be solved with no brute forcing", func(t *testing.T) {
		data := [][]int{
			{0, 0, 0, 2, 6, 0, 7, 0, 1},
			{6, 8, 0, 0, 7, 0, 0, 9, 0},
			{1, 9, 0, 0, 0, 4, 5, 0, 0},
			{8, 2, 0, 1, 0, 0, 0, 4, 0},
			{0, 0, 4, 6, 0, 2, 9, 0, 0},
			{0, 5, 0, 0, 0, 3, 0, 2, 8},
			{0, 0, 9, 3, 0, 0, 0, 7, 4},
			{0, 4, 0, 0, 5, 0, 0, 3, 6},
			{7, 0, 3, 0, 1, 8, 0, 0, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)

		expected := [][]int{
			{4, 3, 5, 2, 6, 9, 7, 8, 1},
			{6, 8, 2, 5, 7, 1, 4, 9, 3},
			{1, 9, 7, 8, 3, 4, 5, 6, 2},
			{8, 2, 6, 1, 9, 5, 3, 4, 7},
			{3, 7, 4, 6, 8, 2, 9, 1, 5},
			{9, 5, 1, 7, 4, 3, 6, 2, 8},
			{5, 1, 9, 3, 2, 6, 8, 7, 4},
			{2, 4, 8, 9, 5, 7, 1, 3, 6},
			{7, 6, 3, 4, 1, 8, 2, 5, 9},
		}
		assert.Equal(t, expected, result)
	})
	t.Run("DFS: Medium Sudoku", func(t *testing.T) {
		data := [][]int{
			{0, 2, 6, 0, 3, 0, 0, 0, 8},
			{9, 0, 0, 6, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 9, 0, 4, 0},
			{0, 0, 7, 3, 0, 2, 0, 0, 0},
			{0, 0, 4, 0, 7, 0, 8, 0, 0},
			{0, 0, 0, 8, 0, 6, 7, 0, 0},
			{0, 5, 0, 7, 2, 0, 0, 0, 0},
			{0, 0, 9, 0, 0, 5, 0, 0, 4},
			{4, 0, 0, 0, 6, 0, 2, 1, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)
		expected := [][]int{
			{1, 2, 6, 4, 3, 7, 5, 9, 8},
			{9, 4, 3, 6, 5, 8, 1, 2, 7},
			{7, 8, 5, 2, 1, 9, 3, 4, 6},
			{8, 6, 7, 3, 9, 2, 4, 5, 1},
			{3, 9, 4, 5, 7, 1, 8, 6, 2},
			{5, 1, 2, 8, 4, 6, 7, 3, 9},
			{6, 5, 1, 7, 2, 4, 9, 8, 3},
			{2, 3, 9, 1, 8, 5, 6, 7, 4},
			{4, 7, 8, 9, 6, 3, 2, 1, 5},
		}
		assert.Equal(t, expected, result)
	})
	t.Run("DFS: Hard Sudoku", func(t *testing.T) {
		data := [][]int{
			{0, 0, 6, 3, 0, 7, 0, 0, 0},
			{0, 0, 4, 0, 0, 0, 0, 0, 5},
			{1, 0, 0, 0, 0, 6, 0, 8, 2},
			{2, 0, 5, 0, 3, 0, 1, 0, 6},
			{0, 0, 0, 2, 0, 0, 3, 0, 0},
			{9, 0, 0, 0, 7, 0, 0, 0, 4},
			{0, 5, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 8, 1, 0, 9, 0, 4, 0},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)
		expected := [][]int{
			{5, 8, 6, 3, 2, 7, 4, 9, 1},
			{7, 2, 4, 8, 9, 1, 6, 3, 5},
			{1, 9, 3, 5, 4, 6, 7, 8, 2},
			{2, 4, 5, 9, 3, 8, 1, 7, 6},
			{8, 6, 7, 2, 1, 4, 3, 5, 9},
			{9, 3, 1, 6, 7, 5, 8, 2, 4},
			{4, 5, 2, 7, 6, 3, 9, 1, 8},
			{3, 1, 9, 4, 8, 2, 5, 6, 7},
			{6, 7, 8, 1, 5, 9, 2, 4, 3},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("DFS: Handles an empty Sudoku starting state", func(t *testing.T) {
		data := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)

		expected := [][]int{
			{9, 8, 7, 6, 5, 4, 3, 2, 1},
			{6, 5, 4, 3, 2, 1, 9, 8, 7},
			{3, 2, 1, 9, 8, 7, 6, 5, 4},
			{8, 7, 9, 4, 3, 6, 2, 1, 5},
			{2, 3, 5, 1, 9, 8, 7, 4, 6},
			{4, 1, 6, 5, 7, 2, 8, 9, 3},
			{7, 9, 3, 8, 4, 5, 1, 6, 2},
			{5, 6, 2, 7, 1, 9, 4, 3, 8},
			{1, 4, 8, 2, 6, 3, 5, 7, 9},
		}

		assert.Equal(t, expected, result)
	})
	t.Run("DFS: Handles an impossible sudoku with no solution", func(t *testing.T) {
		data := [][]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 2},
			{0, 1, 0, 0, 0, 0, 0, 2, 0},
			{0, 0, 1, 0, 0, 0, 2, 0, 0},
			{0, 0, 0, 1, 0, 2, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 1, 0, 0, 0},
			{0, 0, 2, 0, 0, 0, 1, 0, 0},
			{0, 2, 0, 0, 0, 0, 0, 1, 0},
			{2, 0, 0, 0, 0, 0, 0, 0, 1},
		}

		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)

		expected := [][]int(nil)

		assert.Equal(t, expected, result)
	})

	t.Run("DFS: Ian G's sudoku", func(t *testing.T) {
		data := [][]int{
			{1, 0, 0, 4, 3, 0, 0, 0, 0},
			{0, 2, 0, 0, 0, 6, 7, 0, 0},
			{5, 3, 0, 0, 0, 0, 0, 2, 0},
			{4, 6, 0, 0, 0, 7, 5, 0, 0},
			{0, 0, 0, 6, 0, 5, 0, 0, 0},
			{0, 0, 7, 2, 0, 0, 0, 6, 4},
			{0, 4, 0, 0, 0, 0, 0, 5, 7},
			{0, 0, 5, 1, 0, 0, 0, 3, 0},
			{0, 0, 0, 0, 7, 9, 0, 0, 6},
		}
		// create the first node/puzzle
		puzzle := sudoku_solver.Puzzle{}
		puzzle.SetValue(data)

		result := sudoku_solver.Depth_first_search(&puzzle)
		expected := [][]int{
			{1, 7, 9, 4, 3, 2, 6, 8, 5},
			{8, 2, 4, 9, 5, 6, 7, 1, 3},
			{5, 3, 6, 7, 1, 8, 4, 2, 9},
			{4, 6, 2, 3, 8, 7, 5, 9, 1},
			{9, 1, 8, 6, 4, 5, 3, 7, 2},
			{3, 5, 7, 2, 9, 1, 8, 6, 4},
			{6, 4, 1, 8, 2, 3, 9, 5, 7},
			{7, 9, 5, 1, 6, 4, 2, 3, 8},
			{2, 8, 3, 5, 7, 9, 1, 4, 6},
		}
		assert.Equal(t, expected, result)
	})
}
