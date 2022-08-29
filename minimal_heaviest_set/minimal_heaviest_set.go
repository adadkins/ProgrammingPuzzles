package minimal_heaviest_set

import "sort"

func minimalHeaviestSetA(arr []int32) []int32 {
	// Keep adding heaviest item until A > B
	returnArray := []int32{}

	// sort array so we can keep getting the biggest one
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sumArr := 0
	for _, v := range arr {
		sumArr = sumArr + int(v)
	}
	sumReturnArray := 0

	for i := len(arr) - 1; i > 0; i-- {
		// add the biggest one to the return array
		returnArray = append(returnArray, arr[i])
		sumArr = sumArr - int(arr[i])
		sumReturnArray = sumReturnArray + int(arr[i])

		// check if return array is bigger than arr
		if sumReturnArray > sumArr {
			sort.SliceStable(returnArray, func(i, j int) bool {
				return returnArray[i] < returnArray[j]
			})
			return returnArray
		}
	}

	return returnArray
}
