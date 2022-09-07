package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockMaxProfit(t *testing.T) {
	t.Run("Test given example in problem", func(t *testing.T) {
		result := StockMaxProfit([]int{9, 11, 8, 5, 7, 10})
		assert.Equal(t, 5, result)
	})

	t.Run("Test where max profit is very first and last values", func(t *testing.T) {
		result := StockMaxProfit([]int{1, 11, 8, 5, 7, 101})
		assert.Equal(t, 100, result)
	})

	t.Run("Test where there is no profit possible", func(t *testing.T) {
		result := StockMaxProfit([]int{100, 90, 80, 70, 25})
		assert.Equal(t, 0, result)
	})

	t.Run("Test where there is multiple of same profit window", func(t *testing.T) {
		result := StockMaxProfit([]int{50, 100, 10, 60, 5, 55})
		assert.Equal(t, 50, result)
	})
}
