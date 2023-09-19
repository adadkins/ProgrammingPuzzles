package roman_to_integer_test

import (
	"programmingpuzzles/roman_to_integer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		expected := 1
		result := roman_to_integer.RomanToInt("I")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 3", func(t *testing.T) {
		expected := 3
		result := roman_to_integer.RomanToInt("III")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 4", func(t *testing.T) {
		expected := 4
		result := roman_to_integer.RomanToInt("IV")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 5", func(t *testing.T) {
		expected := 5
		result := roman_to_integer.RomanToInt("V")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 9", func(t *testing.T) {
		expected := 9
		result := roman_to_integer.RomanToInt("IX")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 10", func(t *testing.T) {
		expected := 10
		result := roman_to_integer.RomanToInt("X")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 11", func(t *testing.T) {
		expected := 11
		result := roman_to_integer.RomanToInt("XI")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 14", func(t *testing.T) {
		expected := 14
		result := roman_to_integer.RomanToInt("XIV")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 19", func(t *testing.T) {
		expected := 19
		result := roman_to_integer.RomanToInt("XIX")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 30", func(t *testing.T) {
		expected := 30
		result := roman_to_integer.RomanToInt("XXX")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 40", func(t *testing.T) {
		expected := 40
		result := roman_to_integer.RomanToInt("XL")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 44", func(t *testing.T) {
		expected := 44
		result := roman_to_integer.RomanToInt("XLIV")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 70", func(t *testing.T) {
		expected := 70
		result := roman_to_integer.RomanToInt("LXX")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 74", func(t *testing.T) {
		expected := 74
		result := roman_to_integer.RomanToInt("LXXIV")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 91", func(t *testing.T) {
		expected := 91
		result := roman_to_integer.RomanToInt("XCI")
		assert.Equal(t, expected, result)
	})

	t.Run("Test 444", func(t *testing.T) {
		expected := 444
		result := roman_to_integer.RomanToInt("CDXLIV")
		assert.Equal(t, expected, result)
	})
	t.Run("Test 444", func(t *testing.T) {
		expected := 499
		result := roman_to_integer.RomanToInt("CDXLIX")
		assert.Equal(t, expected, result)
	})
}
