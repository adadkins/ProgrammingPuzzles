package programmingpuzzles

import "fmt"

func CeasarCipher(s string, i int) string {
	return caesarCipher(s, int32(i))
}

func caesarCipher(s string, k int32) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	returnString := ""
	for _, v := range s {
		if v == rune('-') {
			returnString = fmt.Sprintf("%v%v", returnString, "-")
			continue
		}
		if v == rune(' ') {
			returnString = fmt.Sprintf("%v%v", returnString, " ")
			continue
		}

		isLetter := false
		//get the index of v
		for kk, vv := range letters {
			if vv == v {
				shiftedLetter := (kk + int(k)) % 26
				if kk > 25 {
					shiftedLetter = shiftedLetter + 26
				}

				returnString = fmt.Sprintf("%v%v", returnString, string(letters[shiftedLetter]))
				isLetter = true
				continue
			}
		}
		if isLetter == false {
			returnString = fmt.Sprintf("%v%v", returnString, string(v))

		}
	}

	return returnString
}
