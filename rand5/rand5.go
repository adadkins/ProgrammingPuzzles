package rand5

import "math/rand"

// this implementation has a tiny tiny tiny chance that it simply never returns. rand7 possibly COULD randomly be >5 a billion times in a row.
func Rand5() int {
	result := rand7()
	if result <= 5 && result >= 1 {
		return result
	}
	return Rand5()
}

func rand7() int {
	return rand.Int() % 7
}

// call random number X number of times and sum them, then take modulus 5. I believe this approaches a random number as x goes to infinity
func Rand5Adder() int {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum = sum + rand7()
	}
	return (sum % 5) + 1
}
