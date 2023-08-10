package sudoku_solver

import "sort"

type Puzzle struct {
	puzzle [][]int
}

func (n *Puzzle) SetValue(input [][]int) {
	n.puzzle = input
}

// breadth first search will pull off the bottom of the data structure
func Breadth_first_search(start *Puzzle) [][]int {
	fillInSquaresWithOnePossibleAnswer(&start.puzzle)

	queue := make([]*Puzzle, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		// get the bottom puzzle and update the queue
		currentNode := queue[0]
		queue = queue[1:]

		// if current puzzle is completed, return it
		if isPuzzleFilled(&currentNode.puzzle) == true {
			return currentNode.puzzle
		}

		var keyValuePairs []struct {
			Row            int
			Column         int
			PossibleValues *[]int
		}
		deadEnd := false
		for r := 0; r < 9 && deadEnd == false; r++ { //loopthrouh the rows
			for c := 0; c < 9 && deadEnd == false; c++ { //loop through the columns
				if currentNode.puzzle[r][c] == 0 {

					// caluclate possible numbers for this square
					possibleNumbers := squarePossibleNumbers(currentNode.puzzle, r, c)

					// If this square has no possible solutions we can return since this path is a dead end
					if len(*possibleNumbers) == 0 {
						deadEnd = true
						break
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

		if deadEnd == false {
			// sort the possible numbers map so we can start with the square with the fewest possible options. This should improve speed
			sort.SliceStable(keyValuePairs, func(i, j int) bool {
				return len((*keyValuePairs[i].PossibleValues)) < len((*keyValuePairs[j].PossibleValues))
			})

			if len(keyValuePairs) > 0 {
				for _, possibleSquareValue := range *keyValuePairs[0].PossibleValues {
					// we need to deep copy here since we're modifying it, and don't want to modify the other version in other other paths
					newPuzzle := deepCopy(currentNode.puzzle)

					// fill in the possible value
					newPuzzle[keyValuePairs[0].Row][keyValuePairs[0].Column] = possibleSquareValue

					// loop over and fill in all the squares that have one possiblity
					fillInSquaresWithOnePossibleAnswer(&newPuzzle)

					if isPuzzleFilled(&newPuzzle) {
						return newPuzzle
					}

					newNode := Puzzle{
						puzzle: newPuzzle,
					}

					// add new puzzle to data structure
					queue = append(queue, &newNode)
				}
			}
		}
	}
	return nil
}

// depth first search will pull off the top of the data structure
func Depth_first_search(start *Puzzle) [][]int {
	fillInSquaresWithOnePossibleAnswer(&start.puzzle)

	queue := make([]*Puzzle, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		// get the top puzzle and update the queue
		currentNode := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		// if current puzzle is completed, return it
		if isPuzzleFilled(&currentNode.puzzle) == true {
			return currentNode.puzzle
		}

		var keyValuePairs []struct {
			Row            int
			Column         int
			PossibleValues *[]int
		}
		deadEnd := false
		for r := 0; r < 9 && deadEnd == false; r++ { //loopthrouh the rows
			for c := 0; c < 9 && deadEnd == false; c++ { //loop through the columns
				if currentNode.puzzle[r][c] == 0 {

					// caluclate possible numbers for this square
					possibleNumbers := squarePossibleNumbers(currentNode.puzzle, r, c)

					// If this square has no possible solutions we can return since this path is a dead end
					if len(*possibleNumbers) == 0 {
						deadEnd = true
						break
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

		if deadEnd == false {
			// sort the possible numbers map so we can start with the square with the fewest possible options. This should improve speed
			sort.SliceStable(keyValuePairs, func(i, j int) bool {
				return len((*keyValuePairs[i].PossibleValues)) < len((*keyValuePairs[j].PossibleValues))
			})

			if len(keyValuePairs) > 0 {
				for _, possibleSquareValue := range *keyValuePairs[0].PossibleValues {
					// we need to deep copy here since we're modifying it, and don't want to modify the other version in other other paths
					newPuzzle := deepCopy(currentNode.puzzle)

					// fill in the possible value
					newPuzzle[keyValuePairs[0].Row][keyValuePairs[0].Column] = possibleSquareValue

					// loop over and fill in all the squares that have one possiblity
					fillInSquaresWithOnePossibleAnswer(&newPuzzle)

					if isPuzzleFilled(&newPuzzle) {
						return newPuzzle
					}

					newNode := Puzzle{
						puzzle: newPuzzle,
					}

					// add new puzzle to data structure
					queue = append(queue, &newNode)
				}
			}
		}
	}
	return nil
}
