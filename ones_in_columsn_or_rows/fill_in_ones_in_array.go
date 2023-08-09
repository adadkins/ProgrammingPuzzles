package fill_in_ones_in_array

func Fill_in_ones(input [][]int) [][]int {
	// loop over and find all the ones, keep track which columns and rows have 1's
	rowsTracker := []int{}
	columnsTracker := []int{}
	for r := 0; r < len(input); r++ { // loop over the rows
		for c := 0; c < len(input[0]); c++ { // loop over the columns
			// if this square has a 1, add to the trackers
			if input[r][c] == 1 {
				rowsTracker = append(rowsTracker, r)
				columnsTracker = append(columnsTracker, c)
			}
		}
	}
	// loop over the rows and fill in
	for _, r := range rowsTracker {
		for c := 0; c < len(input[0]); c++ {
			input[r][c] = 1
		}
	}
	// loop over the columns and fill in
	for _, c := range columnsTracker {
		for r := 0; r < len(input); r++ {
			input[r][c] = 1
		}
	}
	return input
}
