package rand5_test

import (
	"programmingpuzzles/rand5"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand5(t *testing.T) {
	// i cant think of a good way to test randomness that wouldnt occasionally fail
	t.Run("Test recursive Rand5()", func(t *testing.T) {
		mapper := map[int]int{}
		for i := 0; i < 10000; i++ {
			result := rand5.Rand5()
			mapper[result] = mapper[result] + 1
		}

		// verify 1 <= n <= 5
		for k, _ := range mapper {
			assert.LessOrEqual(t, k, 5)
			assert.GreaterOrEqual(t, k, 1)
		}
	})

	// I cant think of a good way to test randomness that wouldnt occasionally fail
	t.Run("Test better implementation Rand5()", func(t *testing.T) {
		mapper := map[int]int{}
		for i := 0; i < 10000; i++ {
			result := rand5.Rand5Adder()
			mapper[result] = mapper[result] + 1
		}

		// verify 1 <= n <= 5
		for k, _ := range mapper {
			assert.LessOrEqual(t, k, 5)
			assert.GreaterOrEqual(t, k, 1)
		}
	})
}
