package binary_flipped_bits

import (
	"fmt"
	"math"
)

func binaryFlippedBits(x int) int {
	counter := 0
	for _, v := range fmt.Sprintf("%b", x) {
		if v == '1' {
			counter++
		}
	}
	return counter
}

// Using the built in string formatter, convert to string representation of binary, loop over string and sum all 1's
func BinaryFlippedBitsUpToN1(input int) int {
	counter := 0
	for i := 0; i <= input; i++ {
		counter = counter + binaryFlippedBits(i)
	}
	return counter
}

// Calculate the flipped bits starting at 1 and store into a map.
// When we get to 5, then 5 is 4+1 so we can just calculate the bits for 5 by pulling the bits for 1 out of the map.
func BinaryFlippedBitsUpToN(input int) int {
	counter := 0
	calculatedValues := map[int]int{}

	calculatedValues[0] = 0

	for i := 1; i <= input; i++ {
		// calculate what the value is removing the first 2's digit, since we have calculated the other exponents
		remainder := i - int(math.Pow(2, float64(int(math.Log2(float64(i))))))

		// since we already calculated the number of bits we can just pull that value from the map
		calculatedValues[i] = 1 + calculatedValues[remainder]
		counter = counter + 1 + calculatedValues[remainder]
	}

	return counter
}
