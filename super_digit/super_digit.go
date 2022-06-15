package super_digit

import (
	"strconv"
)

func SuperDigit(n string, k int) int {
	sum := 0
	for _, kk := range n {
		number := int(kk - '0')
		sum = sum + number
	}
	// convert to string and multiple by concat amount
	ss := strconv.Itoa(sum * k)

	for len(ss) > 1 {
		sum = 0
		for _, kk := range ss {
			number := int(kk - '0')
			sum = sum + number
		}
		ss = strconv.Itoa(sum)
	}
	return int(ss[0] - '0')
}
