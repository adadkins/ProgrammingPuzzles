package game_of_life

func Game_of_life(input *[][]int) *[][]int {
	returnMatrix := make([][]int, len(*input))
	for k := range *input {
		returnMatrix[k] = make([]int, len((*input)[k]))
		copy(returnMatrix[k], (*input)[k])
	}

	// loop over all the squares
	for i := 0; i < len(*input); i++ {
		for j := 0; j < len((*input)[0]); j++ {
			neighborCount := 0
			// check the 8 neighbors

			// right top
			if (i != 0 && j != len((*input)[0])-1) && (*input)[i-1][j+1] == 1 {
				neighborCount = neighborCount + 1
			}
			// right middle
			if j != len((*input)[0])-1 && (*input)[i][j+1] == 1 {
				neighborCount = neighborCount + 1
			}
			// right bottom
			if (i != len(*input)-1 && j != len((*input)[0])-1) && (*input)[i+1][j+1] == 1 {
				neighborCount = neighborCount + 1
			}
			// middle top
			if i != 0 && (*input)[i-1][j] == 1 {
				neighborCount = neighborCount + 1
			}
			// middle bottom
			if i != len(*input)-1 && (*input)[i+1][j] == 1 {
				neighborCount = neighborCount + 1
			}
			// left top
			if (i != 0 && j != 0) && (*input)[i-1][j-1] == 1 {
				neighborCount = neighborCount + 1
			}
			// left middle
			if j != 0 && (*input)[i][j-1] == 1 {
				neighborCount = neighborCount + 1
			}
			// left bottom
			if (i != len(*input)-1 && j != 0) && (*input)[i+1][j-1] == 1 {
				neighborCount = neighborCount + 1
			}

			// update value of square following game of life rules
			if (*input)[i][j] == 1 && (neighborCount < 2 || neighborCount > 3) {
				returnMatrix[i][j] = 0
				continue
			}
			if (*input)[i][j] == 0 && neighborCount == 3 {
				returnMatrix[i][j] = 1
			}
		}
	}
	return &returnMatrix
}

// check if board reaches a steady state ie unchanged after 1000 iterations.
func Does_Board_Reach_Steady_State(input *[][]int) bool {
	for i := 1000; i > 0; i++ {
		//get next stage
		next := Game_of_life(input)

		//check if they're equal
		for i := range *next {
			for j := range (*next)[0] {
				if (*next)[i][j] == (*input)[i][j] {
					return true
				}
			}
		}
	}

	return false
}
