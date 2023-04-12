package sum_numbers_test

import (
	sum_numbers "programmingpuzzles/summing_50_digit_numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumNumbers(t *testing.T) {
	t.Run("First example: sum 5 numbers", func(t *testing.T) {
		input := []string{
			"37107287533902102798797998220837590246510135740250",
			"46376937677490009712648124896970078050417018260538",
			"74324986199524741059474233309513058123726617309629",
			"91942213363574161572522430563301811072406154908250",
			"23067588207539346171171980310421047513778063246676",
		}

		result := sum_numbers.SumNumbers(input)
		assert.Equal(t, "2728190129", result)
	})
	t.Run("First example: sum 6 numbers", func(t *testing.T) {
		input := []string{
			"37107287533902102798797998220837590246510135740250",
			"46376937677490009712648124896970078050417018260538",
			"74324986199524741059474233309513058123726617309629",
			"91942213363574161572522430563301811072406154908250",
			"23067588207539346171171980310421047513778063246676",
			"11111111000000000000000000000000000000000000000000",
		}

		result := sum_numbers.SumNumbers(input)
		assert.Equal(t, "2839301239", result)
	})
	t.Run("First example: sum 1 numbers", func(t *testing.T) {
		input := []string{
			"11111111110000000000000000000000000000000000000000",
		}

		result := sum_numbers.SumNumbers(input)
		assert.Equal(t, "1111111111", result)
	})
	t.Run("First example: sum 2 numbers", func(t *testing.T) {
		input := []string{
			"11111111110000000000000000000000000000000000000000",
			"22222222222222222222222222222222222222222222222222",
		}

		result := sum_numbers.SumNumbers(input)
		assert.Equal(t, "3333333333", result)
	})
	t.Run("First example: sum 3 numbers", func(t *testing.T) {
		input := []string{
			"01111111111111111111111111111111111111111111111111",
			"22222222222222222222222222222222222222222222222222",
			"77777777777777777777777777777777777777777777777777",
		}

		result := sum_numbers.SumNumbers(input)
		assert.Equal(t, "1011111111", result)
	})
}
