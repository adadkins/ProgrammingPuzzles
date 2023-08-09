package sudoku_solver

import (
	"sort"
)

func Solve(puzzle [][]int) [][]int {
	// create a map of possible values so we can sort them
	var keyValuePairs []struct {
		Row            int
		Column         int
		PossibleValues *[]int
	}
	// loop over and fill in all the squares that have one possiblity
	fillInSquaresWithOnePossibleAnswer(&puzzle)

	if isPuzzleFilled(puzzle) {
		return puzzle
	}

	for r := 0; r < 9; r++ { //loopthrouh the rows
		for c := 0; c < 9; c++ { //loop through the columns
			if puzzle[r][c] == 0 {

				// caluclate possible numbers for this square
				possibleNumbers := squarePossibleNumbers(puzzle, r, c)

				// If this square has no possible solutions we can return since this path is a dead end
				if len(*possibleNumbers) == 0 {
					return nil
				}

				// add possible answers to the map so we can iterate and brute force it
				if len(*possibleNumbers) > 1 {
					keyValuePairs = append(keyValuePairs, struct {
						Row            int
						Column         int
						PossibleValues *[]int
					}{Row: r, Column: c, PossibleValues: possibleNumbers})
				}
			}
		}
	}

	// sort the possible numbers map so we can start with the square with the fewest possible options. This should improve speed
	sort.SliceStable(keyValuePairs, func(i, j int) bool {
		return len((*keyValuePairs[i].PossibleValues)) < len((*keyValuePairs[j].PossibleValues))
	})

	if len(keyValuePairs) > 0 {
		for _, possibleSquareValue := range *keyValuePairs[0].PossibleValues {

			// we need to deep copy here since we're modifying it, and don't want to modify the other version in other other paths
			filledInAnswer := deepCopy(puzzle)

			// fill in the possible value
			filledInAnswer[keyValuePairs[0].Row][keyValuePairs[0].Column] = possibleSquareValue

			// try to solve with the guess
			result := Solve(filledInAnswer)
			if result != nil {

				// if the entire square is filled in, we can return it
				if isPuzzleFilled(result) {
					return result
				}
			}
		}
	}

	// if we wanted to use go routines

	// resultChannel := make(chan [][]int)
	// if len(keyValuePairs) > 0 {
	// 	for _, possibleSquareValue := range keyValuePairs[0].PossibleValues {
	// 		go func(val int) {
	// 			// we need to deep copy here since we're modifying it, and don't want to modify the other version in other other paths
	// 			filledInAnswer := deepCopy(answer)

	// 			// fill in the possible value
	// 			filledInAnswer[keyValuePairs[0].Row][keyValuePairs[0].Column] = val

	// 			// try to solve with the guess
	// 			result := Solve(filledInAnswer)
	// 			if result != nil && isPuzzleFilled(result) {
	// 				resultChannel <- result
	// 			}
	// 		}(possibleSquareValue)
	// 	}

	// 	for range keyValuePairs[0].PossibleValues {
	// 		result := <-resultChannel
	// 		if result != nil {
	// 			return result
	// 		}
	// 	}
	// }

	// current branch has no solution
	return nil
}

func isPuzzleFilled(puzzle [][]int) bool {
	for r := 0; r < 9; r++ { //loopthrouh the rows
		for c := 0; c < 9; c++ { //loop through the columns
			if puzzle[r][c] == 0 {
				return false
			}
		}
	}
	return true
}

func squarePossibleNumbers(puzzle [][]int, row, column int) *[]int {
	possibleNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// check all the numbers in the row (loop over the column)
	for c := 0; c < 9; c++ {
		// if the row has a number remove it from possible numbers list
		if puzzle[row][c] != 0 {
			possibleNumbers[puzzle[row][c]-1] = 0
		}
	}

	// check all the numbers in the columns(loop over the row)
	for r := 0; r < 9; r++ {
		// if the row has a number remove it from possible numbers list
		if puzzle[r][column] != 0 {
			possibleNumbers[puzzle[r][column]-1] = 0
		}
	}
	//row is like y corrdinate
	//column is like x coorodinate

	// Check the 3x3 box
	boxStartRow, boxStartCol := row/3*3, column/3*3
	for r := boxStartRow; r < boxStartRow+3; r++ {
		for c := boxStartCol; c < boxStartCol+3; c++ {
			if puzzle[r][c] != 0 {
				possibleNumbers[puzzle[r][c]-1] = 0
			}
		}
	}

	// filter out the 0's
	returnSlice := []int{}
	for _, v := range possibleNumbers {
		if v != 0 {
			returnSlice = append(returnSlice, v)
		}
	}

	return &returnSlice
}

// loop through and fill in all squares that have one possible value
func fillInSquaresWithOnePossibleAnswer(puzzle *[][]int) {
	for {
		modified := false
		for r := 0; r < 9; r++ { //loopthrouh the rows
			for c := 0; c < 9; c++ { //loop through the columns
				if (*puzzle)[r][c] == 0 {
					// caluclate possible numbers for this square
					possibleNumbers := squarePossibleNumbers((*puzzle), r, c)

					// if we only have one possible value for the square, set it then reset our loops
					if len((*possibleNumbers)) == 1 {
						(*puzzle)[r][c] = (*possibleNumbers)[0]
						modified = true
					}
				}
			}
		}

		if modified == false {
			break
		}
	}
}

// copy the puzzle
func deepCopy(input [][]int) [][]int {
	copyGrid := make([][]int, len(input))
	for i := range input {
		copyGrid[i] = make([]int, len(input[i]))
		copy(copyGrid[i], input[i])
	}
	return copyGrid
}
