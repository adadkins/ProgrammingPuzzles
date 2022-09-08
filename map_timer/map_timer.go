package programmingpuzzles

import (
	"sort"

	"golang.org/x/exp/maps"
)

type TimerMap struct {
	mapping map[int]map[int]int
}

func NewTimerMap() *TimerMap {
	return &TimerMap{mapping: make(map[int]map[int]int)}
}

func (m *TimerMap) Set(key, value, time int) {
	mapValue, okay := m.mapping[key]

	// if key exist update the values
	if okay {
		mapValue[time] = value
		//check if any exist after time and change them
		for k := range mapValue {
			if k > time {
				mapValue[k] = value
			}
		}
		return
	}

	// if  key doesnt exist, create map and add to struc mapping
	array := map[int]int{}
	array[time] = value
	m.mapping[key] = array
}

func (m *TimerMap) Get(key, time int) int {
	// get map from mapping
	timeValues, okay := m.mapping[key]
	if !okay {
		return 0
	}
	//check if the exact time value exists
	value1, okay1 := timeValues[time]
	if okay1 {
		return value1
	}

	// check if value was set at a previous time

	//get keys from map and sort them by descending time
	times := maps.Keys(timeValues)
	sort.SliceStable(times, func(i, j int) bool {
		return times[i] < times[j]
	})

	//loop over the time and if was set before time parameter get it. We want the most recent(biggest) time set
	returnValue := 0
	for _, v := range times {
		if v < time {
			returnValue = timeValues[v]
		}
		// dont need to loop over if times are set after time parameter
		if v > time {
			break
		}
	}

	return returnValue
}
