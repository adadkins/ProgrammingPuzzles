package programmingpuzzles

import (
	"fmt"
	"sort"
)

func SmallestBoatCalculator(weights []int, boatLimit int) (boatCounter int) {
	//lets sort weights by descending order
	sort.SliceStable(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	j := 0
	//start with heaviest person, keep adding people until boat limit reached
	for i := len(weights) - 1; i >= j; i-- {
		// if current heaviest person heavier than boat limit, skip him since the boat will sink
		if weights[i] > boatLimit {
			continue
		}

		newBoatWeight := weights[i]
		for jj := j; jj <= i; jj++ {
			if newBoatWeight+weights[jj] <= boatLimit {
				newBoatWeight = newBoatWeight + weights[jj]
				if jj != i {
					continue
				}
			}
			// if we cant add any more people to the boat, move to the next boat, and increase boat counter by 1
			boatCounter++
			//move jj to j so we dont start with the lightest people that have already been put on a boat in next inner loop
			j = jj
			break
		}
	}

	return boatCounter
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
