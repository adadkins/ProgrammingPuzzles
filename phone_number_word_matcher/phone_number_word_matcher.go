package programmingpuzzles

import (
	"fmt"
	"strings"
)

func PhoneWordMatcher(word string, numbers []int) []int {
	returnArray := []int{}
	wordToPhoneNumber := convertWordToPhoneNumber(strings.ToLower(word))
	for _, v := range numbers {
		if strings.Contains(fmt.Sprintf("%v", v), wordToPhoneNumber) {
			returnArray = append(returnArray, v)
		}
	}
	return returnArray
}

func convertWordToPhoneNumber(word string) string {
	// characters to number mapping
	mapping := map[rune]int{
		'a': 2,
		'b': 2,
		'c': 2,
		'd': 3,
		'e': 3,
		'f': 3,
		'g': 4,
		'h': 4,
		'i': 4,
		'j': 5,
		'k': 5,
		'l': 5,
		'm': 6,
		'n': 6,
		'o': 6,
		'p': 7,
		'q': 7,
		'r': 7,
		's': 7,
		't': 8,
		'u': 8,
		'v': 8,
		'w': 9,
		'x': 9,
		'y': 9,
		'z': 9,
	}
	constructedNumber := ""

	// take word, convert it to number
	for _, v := range word {
		digit, ok := mapping[v]
		if ok {
			constructedNumber = fmt.Sprintf("%v%v", constructedNumber, digit)
		}
	}

	return constructedNumber
}
