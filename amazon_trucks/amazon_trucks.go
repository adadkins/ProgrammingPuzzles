package amazon_trucks

import (
	"sort"
)

func GetTrucksForItems(trucks []int, items []int) []int {
	returnTrucks := []int{}
	truckWeightMapping := map[int]int{}
	// create a mapping of trucks to values
	for k, v := range trucks {
		// if we have duplicates we dont need the largest index truck
		if _, ok := truckWeightMapping[v]; ok == false {
			truckWeightMapping[v] = k
		}
	}

	// add -1 for not found
	truckWeightMapping[-1] = -1

	sort.Slice(trucks, func(i, j int) bool {
		return trucks[i] > trucks[j]
	})

	// keep a map so we dont have to continually search for known values
	itemTruckMapping := map[int]int{}

	// loop over the items and find the smallest truck index
	for itemIndex, itemWeight := range items {
		if val, ok := itemTruckMapping[itemWeight]; ok {
			items[itemIndex] = val
			break
		}
		iterTruckValue := -1

		// loop over the trucks and find the truck with the capacity that is smallest but less than
		for _, truckCapacity := range trucks {
			// check if the item can fit in the truck
			if itemWeight < truckCapacity {
				iterTruckValue = truckCapacity
			}
			if itemWeight >= truckCapacity {
				break
			}
		}

		itemTruckMapping[itemWeight] = truckWeightMapping[iterTruckValue]
		returnTrucks = append(returnTrucks, truckWeightMapping[iterTruckValue])
	}

	return returnTrucks
}

func GetTrucksForItems1(trucks []int, items []int) []int {
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
