package roman_to_integer

import (
	"strings"
)

func RomanToInt(input string) int {

	//mapping := map[int]string{1: "I", 5: "V", 9: "IX", 10: "X", 50: "L", 90: "XC", 100: "C", 500: "D"}

	numerals := []struct {
		digit      int
		conversion string
	}{
		{400, "CD"},
		{40, "XL"},
		{4, "IV"},
		{90, "XC"},
		{9, "IX"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}

	returnInt := 0

	for _, numeral := range numerals {
		//check if the input string contains that character
		// result := strings.Contains(input, numeral.conversion)
		// fmt.Println(result)
		for strings.Contains(input, numeral.conversion) {
			//add the value to the return value
			returnInt = returnInt + numeral.digit

			//remove the numeral from the input string
			input1, input2, _ := strings.Cut(input, numeral.conversion)
			input = input1 + input2
		}
	}
	return returnInt
}
