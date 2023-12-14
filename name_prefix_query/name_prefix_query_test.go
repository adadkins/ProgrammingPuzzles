package name_prefix_query_test

import (
	"programmingpuzzles/name_prefix_query"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamePrefixQuery(t *testing.T) {
	t.Run("Given Example", func(t *testing.T) {
		// given
		queries := []string{"rick"}
		strings := []string{"rick", "rickroll", "rock"}

		// when
		result := name_prefix_query.FindCompletePrefixes(strings, queries)

		// then
		expected := []int{1}
		assert.Equal(t, expected, result)
	})

	t.Run("Empty Inputs", func(t *testing.T) {
		// given
		queries := []string{}
		strings := []string{}

		// when
		result := name_prefix_query.FindCompletePrefixes(strings, queries)

		// then
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("No Matches", func(t *testing.T) {
		// given
		queries := []string{"abc", "def", "ghi", "jkl"}
		strings := []string{"aaaaaa", "bbbbbb", "cccccc", "dddddd", "eeeeee", "fffffff"}
		// when
		result := name_prefix_query.FindCompletePrefixes(strings, queries)

		// then
		expected := []int{0, 0, 0, 0}
		assert.Equal(t, expected, result)
	})

	t.Run("Some Matches", func(t *testing.T) {
		// given
		queries := []string{"a", "ab", "ac", "af"}
		strings := []string{"aaaaaa", "abbbbbb", "acccccc", "adddddd", "aeeeeee", "afffffff"}
		// when
		result := name_prefix_query.FindCompletePrefixes(strings, queries)

		// then
		expected := []int{6, 1, 1, 1}
		assert.Equal(t, expected, result)
	})
}
