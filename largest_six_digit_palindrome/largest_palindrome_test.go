package largest_six_digit_palindrome_test

import (
	"math/rand"
	"programmingpuzzles/largest_six_digit_palindrome"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLargestPalindrome(t *testing.T) {
	t.Run("Largest palindrome for example 101110 is 101101", func(t *testing.T) {
		result := largest_six_digit_palindrome.LargestPalindrome(101110)
		assert.Equal(t, 101101, result)
	})

	t.Run("Largest palindrome for example 800000 is 793397", func(t *testing.T) {
		result := largest_six_digit_palindrome.LargestPalindrome(800000)
		assert.Equal(t, 793397, result)
	})

	t.Run("Largest palindrome for input 1000000 is 906700", func(t *testing.T) {
		result := largest_six_digit_palindrome.LargestPalindrome(906700)
		assert.Equal(t, 906609, result)
	})

	t.Run("Largest palindrome for input 580100 is 580085", func(t *testing.T) {
		result := largest_six_digit_palindrome.LargestPalindrome(580100)
		assert.Equal(t, 580085, result)
	})
}

// benchmark both palindrome functions to see which works better
func BenchmarkLargestPalindrome(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		num := rand.Intn(10000000-100000+1) + 100000
		largest_six_digit_palindrome.IsPalidrome(num)
	}
}

func BenchmarkLargestPalindrome_stringConversion(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		num := rand.Intn(10000000-100000+1) + 100000
		largest_six_digit_palindrome.IsPalidrome_stringConverstion(num)
	}
}
