package closest_points_test

import (
	"programmingpuzzles/closest_points"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Point = closest_points.Point

func TestClosestPoints(t *testing.T) {
	t.Run("Given Example", func(t *testing.T) {
		input := []Point{{1, 1}, {-1, -1}, {3, 4}, {6, 1}, {-1, -6}, {-4, -3}}
		expected1, expected2 := Point{1, 1}, Point{-1, -1}
		result1, result2 := closest_points.ClosestPoints(input)
		assert.Equal(t, expected1, result1)
		assert.Equal(t, expected2, result2)
	})
}
