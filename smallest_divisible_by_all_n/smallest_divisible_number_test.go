package smallest_divisible_test

import (
	smallest_divisible "programmingpuzzles/smallest_divisible_by_all_n"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestDivisibleNumber(t *testing.T) {
	t.Run("Smallest divisible number 10 is 2520", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(10)
		assert.Equal(t, 2520, result)
	})

	t.Run("Smallest divisible number 11 is 27720", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(11)
		assert.Equal(t, 27720, result)
	})

	t.Run("Smallest divisible number 12 is 27720", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(12)
		assert.Equal(t, 27720, result)
	})

	t.Run("Smallest divisible number 13 is 360360", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(13)
		assert.Equal(t, 360360, result)
	})
	t.Run("Smallest divisible number 14 is 360360", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(14)
		assert.Equal(t, 360360, result)
	})

	t.Run("Smallest divisible number 16 is 720720", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(16)
		assert.Equal(t, 720720, result)
	})

	t.Run("Smallest divisible number 17 is 360360", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(17)
		assert.Equal(t, 12252240, result)
	})

	t.Run("Smallest divisible number 19 is 232792560", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(19)
		assert.Equal(t, 232792560, result)
	})

	t.Run("Smallest divisible number 20 is 232792560", func(t *testing.T) {
		result := smallest_divisible.SmallestDivisibleNumber(20)
		assert.Equal(t, 232792560, result)
	})
}

func TestGetPrimes(t *testing.T) {
	t.Run("Primes for 10", func(t *testing.T) {
		result := smallest_divisible.GetAllPrimes(10)
		assert.Equal(t, map[int]int{2: 1, 5: 1}, result)
	})

	t.Run("Primes for 25", func(t *testing.T) {
		result := smallest_divisible.GetAllPrimes(25)
		assert.Equal(t, map[int]int{5: 2}, result)
	})
	t.Run("Primes for 1000", func(t *testing.T) {
		result := smallest_divisible.GetAllPrimes(1000)
		assert.Equal(t, map[int]int{5: 3, 2: 3}, result)
	})

	t.Run("Primes for 121", func(t *testing.T) {
		result := smallest_divisible.GetAllPrimes(121)
		assert.Equal(t, map[int]int{11: 2}, result)
	})
}
