package integer_to_roman_test

import (
	"programmingpuzzles/integer_to_roman"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		expected := "I"
		result := integer_to_roman.IntToRoman(1)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 3", func(t *testing.T) {
		expected := "III"
		result := integer_to_roman.IntToRoman(3)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 4", func(t *testing.T) {
		expected := "IV"
		result := integer_to_roman.IntToRoman(4)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 5", func(t *testing.T) {
		expected := "V"
		result := integer_to_roman.IntToRoman(5)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 9", func(t *testing.T) {
		expected := "IX"
		result := integer_to_roman.IntToRoman(9)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 10", func(t *testing.T) {
		expected := "X"
		result := integer_to_roman.IntToRoman(10)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 11", func(t *testing.T) {
		expected := "XI"
		result := integer_to_roman.IntToRoman(11)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 14", func(t *testing.T) {
		expected := "XIV"
		result := integer_to_roman.IntToRoman(14)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 19", func(t *testing.T) {
		expected := "XIX"
		result := integer_to_roman.IntToRoman(19)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 30", func(t *testing.T) {
		expected := "XXX"
		result := integer_to_roman.IntToRoman(30)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 40", func(t *testing.T) {
		expected := "XL"
		result := integer_to_roman.IntToRoman(40)
		assert.Equal(t, expected, result)
	})
	t.Run("Test 44", func(t *testing.T) {
		expected := "XLIV"
		result := integer_to_roman.IntToRoman(44)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 70", func(t *testing.T) {
		expected := "LXX"
		result := integer_to_roman.IntToRoman(70)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 74", func(t *testing.T) {
		expected := "LXXIV"
		result := integer_to_roman.IntToRoman(74)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 91", func(t *testing.T) {
		expected := "XCI"
		result := integer_to_roman.IntToRoman(91)
		assert.Equal(t, expected, result)
	})

	t.Run("Test 444", func(t *testing.T) {
		expected := "CDXLIV"
		result := integer_to_roman.IntToRoman(444)
		assert.Equal(t, expected, result)
	})
}
