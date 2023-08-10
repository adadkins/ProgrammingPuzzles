package sum_two_numbers

// naive implementation, two loops
func Sum_two_numbers(input []int, sumTarget int) bool {
	for x := 0; x < len(input); x++ {
		for xx := x; xx < len(input); xx++ {
			if input[x]+input[xx] == sumTarget {
				return true
			}
		}
	}
	return false
}

// how to do it in one pass
func Sum_two_numbers_single_pass(input []int, sumTarget int) bool {
	checker := map[int]bool{}
	for _, v := range input {
		if checker[sumTarget-v] == true {
			return true
		}
		checker[v] = true
	}
	return false
}
