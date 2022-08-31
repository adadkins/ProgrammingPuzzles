package interval_overlap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test
func TestIntervalOverlap(t *testing.T) {
	t.Run("Test overlapping intervals", func(t *testing.T) {
		result := IntervalOverlap(Interval{1, 5}, Interval{4, 10})
		assert.Equal(t, true, result)
	})

	t.Run("Test overlapping intervals, largest interval first", func(t *testing.T) {
		result := IntervalOverlap(Interval{4, 10}, Interval{1, 5})
		assert.Equal(t, true, result)
	})

	t.Run("Test overlapping intervals, completley overlapping", func(t *testing.T) {
		result := IntervalOverlap(Interval{1, 10}, Interval{3, 5})
		assert.Equal(t, true, result)
	})
	t.Run("Test overlapping intervals, completley overlapping", func(t *testing.T) {
		result := IntervalOverlap(Interval{40, 60}, Interval{1, 100})
		assert.Equal(t, true, result)
	})

	t.Run("Test non overlapping intervals", func(t *testing.T) {
		result := IntervalOverlap(Interval{1, 5}, Interval{10, 15})
		assert.Equal(t, false, result)
	})

	t.Run("Test non overlapping intervals", func(t *testing.T) {
		result := IntervalOverlap(Interval{40, 60}, Interval{80, 100})
		assert.Equal(t, false, result)
	})
}

func TestIntervalsOverlap(t *testing.T) {
	t.Run("Test overlapping intervals", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{1, 5}, {4, 10}, {9, 15}})
		assert.Equal(t, true, result)
	})

	t.Run("Test overlapping intervals, largest interval first", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{4, 10}, {1, 5}, {0, 1}})
		assert.Equal(t, true, result)
	})

	t.Run("Test overlapping intervals, completley overlapping", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{1, 10}, {3, 5}, {4, 5}})
		assert.Equal(t, true, result)
	})
	t.Run("Test overlapping intervals, completley overlapping", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{40, 60}, {1, 100}, {50, 55}})
		assert.Equal(t, true, result)
	})

	t.Run("Test non overlapping intervals", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{1, 5}, {10, 15}, {20, 25}})
		assert.Equal(t, false, result)
	})

	t.Run("Test non overlapping intervals", func(t *testing.T) {
		result := IntervalsOverlap([]Interval{{40, 60}, {80, 100}, {110, 120}})
		assert.Equal(t, false, result)
	})
}
