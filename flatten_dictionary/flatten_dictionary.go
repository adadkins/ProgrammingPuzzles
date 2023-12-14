package flatten_dictionary

import "fmt"

type Dictionary struct {
	key map[string]interface{}
}

func Flatten(currentKey string, input map[string]interface{}) map[string]int {
	results := map[string]int{}
	for k, v := range input {
		if _, ok := v.(int); ok {
			if currentKey != "" {
				results[fmt.Sprintf("%s.%s", currentKey, k)] = v.(int)
				continue
			}
			results[k] = v.(int)
		}
		if _, ok := v.(map[string]interface{}); ok {
			d := Flatten(k, v.(map[string]interface{}))

			// loop over the returned mapping and append current key to keys
			for kk, vv := range d {
				if currentKey != "" {
					results[fmt.Sprintf("%s.%s", currentKey, kk)] = vv
					continue
				}
				results[kk] = vv
			}
		}
	}
	return results
}
