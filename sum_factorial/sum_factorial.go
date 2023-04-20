package sum_factorial

import (
	"math/big"
)

func SumFactorialArray(input []int) []int {
	returnArray := make([]int, len(input))

	for k, v := range input {
		result := sumDigits(v)
		returnArray[k] = result
	}
	return returnArray
}

func sumDigits(n int) int {
	factorial := new(big.Int).MulRange(1, int64(n))

	digitSum := 0
	for _, digit := range factorial.String() {
		if digit >= '0' && digit <= '9' {
			digitSum += int(digit - '0')
		}
	}
	return digitSum
}
