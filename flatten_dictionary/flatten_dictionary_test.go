package flatten_dictionary_test

import (
	"programmingpuzzles/flatten_dictionary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenDictionary(t *testing.T) {
	t.Run("Given Example", func(t *testing.T) {
		input := map[string]interface{}{
			"key": 3,
			"foo": map[string]interface{}{
				"a": 5,
				"bar": map[string]interface{}{
					"baz": 8,
				},
			},
		}
		result := flatten_dictionary.Flatten("", input)
		expected := map[string]int{
			"key":         3,
			"foo.a":       5,
			"foo.bar.baz": 8,
		}
		assert.Equal(t, expected, result)

	})
}
