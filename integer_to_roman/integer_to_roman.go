package integer_to_roman

func IntToRoman(num int) string {

	//mapping := map[int]string{1: "I", 5: "V", 9: "IX", 10: "X", 50: "L", 90: "XC", 100: "C", 500: "D"}

	numerals := []struct {
		digit      int
		conversion string
	}{
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	returnNumeral := ""

	for _, numeral := range numerals {
		for num >= numeral.digit {
			returnNumeral = returnNumeral + numeral.conversion
			num = num - numeral.digit
		}
	}
	return returnNumeral
}
