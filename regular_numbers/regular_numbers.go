package programmingpuzzles

func RegularNumbers(n int) []int {
	returnArray := []int{1}
	possibleNumber := 2
	//this is basically a prime factors problem where factors are ONLY 2,3, or 5
	for len(returnArray) < n {
		factors := getPrimeFactors(possibleNumber)
		isRegNumber := isRegularNumber(factors)
		if isRegNumber {
			returnArray = append(returnArray, possibleNumber)
		}
		possibleNumber++
	}
	return returnArray
}

func isRegularNumber(n []int) bool {
	for _, v := range n {
		// if we find NOT 2 3 5, we can skip
		if v != 2 && v != 3 && v != 5 {
			return false
		}
	}
	return true
}

func getPrimeFactors(n int) []int {
	factors := []int{}
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			isPrime := true
			//check if factor already exists in keys, we only want PRIME factors
			for _, v := range factors {
				if i%v == 0 {
					isPrime = false
				}
			}
			if isPrime {
				factors = append(factors, i)
			}
		}
	}
	return factors
}
