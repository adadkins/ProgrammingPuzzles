package interval_overlap

import "sort"

type Interval struct {
	start int
	end   int
}

func IntervalsOverlap(arr []Interval) bool {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].start < arr[j].start
	})
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if IntervalOverlap(arr[i], arr[j]) {
				return true
			}
		}
	}

	return false
}

func IntervalOverlap(a, b Interval) bool {
	if a.start >= b.start && a.start < b.end {
		return true
	}
	if b.start >= a.start && b.start < a.end {
		return true
	}
	return false
}
