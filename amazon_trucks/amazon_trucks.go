package amazon_trucks

import (
	"sort"
)

func GetTrucksForItems(trucks []int, items []int) []int {
	// map the truck weights to their index. Later we will want to get the index by truck weight
	truckWeightIndexMapping := map[int]int{}
	for index, weight := range trucks {
		// check if we already have mapped a truck with this weight
		if _, ok := truckWeightIndexMapping[weight]; ok == false {
			truckWeightIndexMapping[weight] = index
		}
	}

	// now that we have the mapping of weight -> index, we can sort the trucks
	sort.SliceStable(trucks, func(i, j int) bool {
		return trucks[i] > trucks[j]
	})

	returnList := []int{}

	// now that trucks are sorted, we can then go through and find the truck that JUST fits each package
	for _, itemWeight := range items {
		foundTruckWeight := -1
		// loop over the trucks
		for _, truckWeight := range trucks {
			if truckWeight > itemWeight {
				foundTruckWeight = truckWeight
			}
			if truckWeight <= itemWeight {
				break
			}
		}

		// now that we know the truck that can carry the package, get the original index
		returnList = append(returnList, truckWeightIndexMapping[foundTruckWeight])

	}
	return returnList
}
