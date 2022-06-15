package grid_challenge

import (
	"sort"
	"strings"
)

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func GridChallenge(grid []string) string {
	// sort each row
	for i, v := range grid {
		strs := strings.Split(v, "")
		sort.Strings(strs)
		grid[i] = string(strings.Join(strs[:], ""))
	}
	// loop over reach item in each column and make sure its rune increasing
	lenString := len(grid[0])
	for col := 0; col < lenString; col++ {
		for row := 0; row < lenString-1; row++ {
			if grid[row][col] > grid[row+1][col] {
				return "NO"
			}
		}
	}
	return "YES"
}
