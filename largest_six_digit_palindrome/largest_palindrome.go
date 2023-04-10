package largest_six_digit_palindrome

import "strconv"

// returns largest six digit palindrome less than n which is a multiple of two 3 digit numbers
func LargestPalindrome(n int) int {
	for n > 10000 {
		n--
		if IsPalidrome(n) {
			if isMultiple(n) {
				return n
			}
		}
	}
	return 0
}

// make public to benchmark
// benchmark shows average of 15 np/op, about 2-3x faster than string conversion which has about 40 ns/op. Not sure worth it for stupid complexity though
func IsPalidrome(n int) bool {
	if (n/100000 == n%10) && ((n/10000)%10 == (n%100)/10) && ((n/1000)%10 == (n%1000)/100) {
		return true
	}
	return false
}

// make public to benchmark
func IsPalidrome_stringConverstion(n int) bool {
	s := strconv.Itoa(n)
	if s[0] == s[5] && s[1] == s[4] && s[2] == s[3] {
		return true
	}
	return false
}

func isMultiple(n int) bool {
	for i := 999; i >= 100; i-- {
		if n%i == 0 && n/i >= 100 && n/i <= 999 {
			return true
		}
		if i*i < n {
			break
		}
	}
	return false
}
