package sum_primes_test

import (
	"programmingpuzzles/sum_primes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumPrimes(t *testing.T) {
	t.Run("Test sum primes 5 is 10 (2+3+5)", func(t *testing.T) {
		result := sum_primes.SumPrimes(5)
		assert.Equal(t, 10, result)
	})

	t.Run("Test sum primes 10 is 17 (2+3+5+7)", func(t *testing.T) {
		result := sum_primes.SumPrimes(10)
		assert.Equal(t, 17, result)
	})

	t.Run("Test sum primes 11 is 28 (2+3+5+7+11)", func(t *testing.T) {
		result := sum_primes.SumPrimes(11)
		assert.Equal(t, 28, result)
	})

	t.Run("Test sum primes 13 is 28 (2+3+5+7+11_13)", func(t *testing.T) {
		result := sum_primes.SumPrimes(13)
		assert.Equal(t, 41, result)
	})

	t.Run("Test sum primes 100 is 1060", func(t *testing.T) {
		result := sum_primes.SumPrimes(100)
		assert.Equal(t, 1060, result)
	})

	t.Run("Test sum primes 1000 is 76127", func(t *testing.T) {
		result := sum_primes.SumPrimes(1000)
		assert.Equal(t, 76127, result)
	})

	t.Run("Test sum primes 100000 is 76127", func(t *testing.T) {
		result := sum_primes.SumPrimes(100000)
		assert.Equal(t, 454396537, result)
	})
}

func TestSumPrimesSieve(t *testing.T) {
	t.Run("Test sum primes 5 is 10 (2+3+5)", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(5)
		assert.Equal(t, 10, result)
	})

	t.Run("Test sum primes 10 is 17 (2+3+5+7)", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(10)
		assert.Equal(t, 17, result)
	})

	t.Run("Test sum primes 11 is 28 (2+3+5+7+11)", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(11)
		assert.Equal(t, 28, result)
	})

	t.Run("Test sum primes 13 is 41 (2+3+5+7+11+13)", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(13)
		assert.Equal(t, 41, result)
	})

	t.Run("Test sum primes 100 is 1060", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(100)
		assert.Equal(t, 1060, result)
	})

	t.Run("Test sum primes 1000 is 76127", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(1000)
		assert.Equal(t, 76127, result)
	})

	t.Run("Test sum primes 100000 is 76127", func(t *testing.T) {
		result := sum_primes.SumPrimesSieve(5000000)
		assert.Equal(t, 838596693108, result)
	})
}

func TestSumPrimesArray(t *testing.T) {
	t.Run("Find sums of two numbers, one prime one not prime", func(t *testing.T) {
		result := sum_primes.SumPrimesArray([]int{5, 10})
		expected := []int{10, 17}
		assert.Equal(t, expected, result)
	})

	t.Run("Find sums of 4 numbers", func(t *testing.T) {
		result := sum_primes.SumPrimesArray([]int{5, 10, 100, 200})
		expected := []int{10, 17, 1060, 4227}
		assert.Equal(t, expected, result)
	})

	t.Run("Find sums of numltiple powers of 10 up to 10,0000", func(t *testing.T) {
		result := sum_primes.SumPrimesArray([]int{10, 100, 1000, 10000})
		expected := []int{17, 1060, 76127, 5736396}
		assert.Equal(t, expected, result)
	})
}
