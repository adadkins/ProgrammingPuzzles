package name_prefix_query

import "strings"

func FindCompletePrefixes(names []string, queries []string) []int {
	// return slice to hold each query results
	queryCounter := []int{}

	// loop over queries
	for _, query := range queries {
		counter := 0
		// loop over the names and see if it starts with the query
		for _, name := range names {
			// make sure name is longer so its actually a substring
			if strings.HasPrefix(name, query) && len(name) > len(query) {
				counter++
			}
		}
		// append results to our return slice
		queryCounter = append(queryCounter, counter)
	}
	return queryCounter
}
