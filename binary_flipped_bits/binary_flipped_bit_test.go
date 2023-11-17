package binary_flipped_bits_test

import (
	"programmingpuzzles/binary_flipped_bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryFlippedBits(t *testing.T) {
	t.Run("Test first 6,000 integers, comparing both methods", func(t *testing.T) {
		input := 10000
		for i := 1; i < input; i++ {
			result := binary_flipped_bits.BinaryFlippedBitsUpToN(i)
			slowResult := binary_flipped_bits.BinaryFlippedBitsUpToN1(i)

			assert.Equal(t, slowResult, result, "Incorrect for %s", i)
		}
	})

	t.Run("Test string method with 1million ", func(t *testing.T) {
		input := 1000000
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 9884999
		assert.Equal(t, expected, result)
	})

	t.Run("Test caching method with 1million ", func(t *testing.T) {
		input := 1000000
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 9884999
		assert.Equal(t, expected, result)
	})

	t.Run("Test string method with 9million ", func(t *testing.T) {
		input := 9000000
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 102844999
		assert.Equal(t, expected, result)
	})

	t.Run("Test caching method with 9 million ", func(t *testing.T) {
		input := 9000000
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 102844999
		assert.Equal(t, expected, result)
	})

	t.Run("Test string method with 15 million and 11", func(t *testing.T) {
		input := 15000011
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 176841804
		assert.Equal(t, expected, result)
	})

	t.Run("Test caching method with 15 million and 11", func(t *testing.T) {
		input := 15000011
		result := binary_flipped_bits.BinaryFlippedBitsUpToN1(input)
		expected := 176841804
		assert.Equal(t, expected, result)
	})
}
