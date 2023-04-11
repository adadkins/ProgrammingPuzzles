package smallest_divisible

func SmallestDivisibleNumber(n int) int {
	returnInt := 1
	allPrimes := map[int]int{}

	for i := 2; i <= n; i++ {
		//get all prime factors
		primes := GetAllPrimes(i)
		// check if they exist in the mapping
		for k, v := range primes {
			currentPrimeCount := allPrimes[k]
			if v > currentPrimeCount {
				//multiply by primes that are not in the mapping
				//update our return value
				returnInt = returnInt * k * (v - currentPrimeCount)

				//update the mapping with those missing primes
				allPrimes[k] = v
			}
		}
	}
	return returnInt
}

// we can store the prime mapping to look up later so we dont have to keep calculating what the factors of 25 are for instance
var storedPrimeMapping = map[int]map[int]int{}

func GetAllPrimes(n int) map[int]int {
	primeMapping := map[int]int{}

	for n > 1 {
		alreadyFactoredPrimes, exists := storedPrimeMapping[n]
		if exists {
			// add the primes remainding to new prime list
			for k, v := range alreadyFactoredPrimes {
				primeMapping[k] = primeMapping[k] + v
			}

			break
		}

		for i := 2; i <= n; i++ {
			if n%i == 0 {
				primeMapping[i] = primeMapping[i] + 1

				//update our numbas, not sure this works
				n = n / i
				i = 2
				break
			}
		}
	}

	storedPrimeMapping[n] = primeMapping

	return primeMapping
}
