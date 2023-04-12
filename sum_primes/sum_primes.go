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
	for k, _ := range *primes {
		if n%k == 0 {
			return false
		}
	}
	(*primes)[n] = true
	return true
}

// use a seive, should be faster as it scales O(n)
func SumPrimesSieve(n int) int {
	values := make([]bool, n+1)
	sum := 2

	for i := 3; i <= n; i = i + 2 {
		if values[i] == true {
			continue
		}
		sum += i

		for ii := i * i; ii <= n; ii = ii + 2*i {
			values[ii] = true
		}
	}

	return sum
}
