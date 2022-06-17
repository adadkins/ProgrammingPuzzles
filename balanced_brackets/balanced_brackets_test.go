package balanced_brackets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridChallenge(t *testing.T) {
	t.Run("((())) trivial result YES", func(t *testing.T) {
		result := isBalanced("((()))")
		assert.Equal(t, "YES", result)
	})

	t.Run("(((}}} trivial result NO", func(t *testing.T) {
		result := isBalanced("(((}}}")
		assert.Equal(t, "NO", result)
	})

	t.Run("is bracket balanced example test case 0", func(t *testing.T) {
		result := isBalanced("{[()]}")
		assert.Equal(t, "YES", result)

		result = isBalanced("{[(])}")
		assert.Equal(t, "NO", result)

		result = isBalanced("{{[[(())]]}}")
		assert.Equal(t, "YES", result)
	})

	t.Run("is bracket balanced example test case 1", func(t *testing.T) {
		result := isBalanced("{{([])}}")
		assert.Equal(t, "YES", result)

		result = isBalanced("{{)[](}}")
		assert.Equal(t, "NO", result)
	})
	t.Run("is bracket balanced example test case 2", func(t *testing.T) {
		result := isBalanced("{(([])[])[]}")
		assert.Equal(t, "YES", result)

		result = isBalanced("{(([])[])[]]}")
		assert.Equal(t, "NO", result)

		result = isBalanced("{(([])[])[]}[]")
		assert.Equal(t, "YES", result)
	})
}
