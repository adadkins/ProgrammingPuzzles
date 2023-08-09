package rotate_array

// modify existing array
func Rotate_array_in_place(input []int, rotation int) []int {
	// loop over each shift, then loop down the line
	for i := 0; i < rotation; i++ {
		temp := input[i]
		for ii := i; ii < len(input); ii = ii + rotation {
			if ii+(rotation+1) > len(input) {
				input[ii] = temp
				continue
			}
			input[ii] = input[ii+rotation]
		}
	}

	return input
}

// dumb way by copying to new array.
func Rotate_array_in_array(input []int, rotation int) []int {
	returnArray := []int{}
	for i := 0; i < len(input); i++ {
		returnArray = append(returnArray, input[(i+rotation)%len(input)])
	}
	return returnArray
}
