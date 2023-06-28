package main

func Longest_array_of_distinct_elements(arr []int) int {
	longestCounter := 0
	mapping := map[int]int{}
	counter := 0
	for i := 0; i < len(arr); i++ {
		_, ok := mapping[arr[i]]
		if !ok {
			mapping[arr[i]] = 1
			counter++
			continue
		}

		//if we stumble into a dupe, store the counter in longest counter
		//reset mapping
		if counter > longestCounter {
			longestCounter = counter
		}
		counter = 1
		mapping = map[int]int{}
		mapping[arr[i]] = 1

	}
	if counter > longestCounter {
		longestCounter = counter
	}
	return longestCounter
}
