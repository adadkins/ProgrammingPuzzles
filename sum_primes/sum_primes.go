package sum_primes

func SumPrimes(n int) int {
	sum := 0
	primes := make(map[int]bool)
	ptr := &primes
	for i := 2; i <= n; i++ {
		if IsPrime(i, ptr) {
			sum = sum + i
		}
	}

	return sum
}

func IsPrime(n int, primes *map[int]bool) bool {
	for k := range *primes {
		if n%k == 0 {
			return false
		}
	}
	(*primes)[n] = true
	return true
}

// use a seive, should be faster as it scales O(n)
func SumPrimesSieve(n int) []bool {
	isNotPrime := make([]bool, n+1)

	for i := 2; i <= n; i = i + 1 {
		if isNotPrime[i] {
			continue
		}

		for ii := i + i; ii <= n; ii = ii + i {
			isNotPrime[ii] = true
		}
	}

	return isNotPrime
}

func SumPrimesArray(input []int) []int {
	biggestN := 0
	// given an array of N's, find the biggest N
	for _, v := range input {
		if v > biggestN {
			biggestN = v
		}
	}

	// now that we have the biggest N, find the sieve. False == prime number
	sieve := SumPrimesSieve(int(biggestN))

	sum := 0
	calculatedSums := make([]int, len(sieve))

	// loop over the sieve and sum up primes. Start at 2 since thats the first prime
	for i := 2; i < len(sieve); i++ {
		if !sieve[i] {
			sum = sum + i
		}
		calculatedSums[i] = sum
	}

	// create an array to store each input N's sum (for hacker rank we would need to print this)
	returnArray := make([]int, len(input))

	// loop over the N's and add the calculated sum from the array. N is the index in the array. value is the sum for N
	for k, v := range input {
		returnArray[k] = calculatedSums[v]
	}
	return returnArray
}
