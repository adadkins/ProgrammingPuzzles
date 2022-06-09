package programmingpuzzles

import "fmt"

func CeasarCipher(s string, i int) string {
	return caesarCipher(s, int32(i))
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	returnString := ""
	for _, v := range s {
		if v == rune('-') {
			returnString = fmt.Sprintf("%v%v", returnString, "-")
			continue
		}
		//get the index of v
		for kk, vv := range letters {
			if vv == v {

				shiftedLetter := (kk + int(k)) % 26
				if kk > 25 {
					shiftedLetter = shiftedLetter + 26
				}

				//handle capitalizatio
				returnString = fmt.Sprintf("%v%v", returnString, string(letters[shiftedLetter]))
				continue
			}
		}
	}

	return returnString
}
