package grid_challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridChallenge(t *testing.T) {
	t.Run("Grid challenge trivial result YES", func(t *testing.T) {
		result := GridChallenge([]string{"abc", "bcd", "cde"})
		assert.Equal(t, "YES", result)
	})

	t.Run("Grid challenge trivial result NO", func(t *testing.T) {
		result := GridChallenge([]string{"cdef", "bcde", "abcd"})
		assert.Equal(t, "NO", result)
	})

	t.Run("Grid challenge example test case 0", func(t *testing.T) {
		result := GridChallenge([]string{"eabcd", "fghij", "olkmn", "trpqs", "xywuv"})
		assert.Equal(t, "YES", result)
	})

	t.Run("Grid challenge example test case 1", func(t *testing.T) {
		result := GridChallenge([]string{"abc", "lmp", "qrt"})
		assert.Equal(t, "YES", result)

		result = GridChallenge([]string{"mpxz", "abcd", "wlmf"})
		assert.Equal(t, "NO", result)

		result = GridChallenge([]string{"abc", "hjk", "mpq", "rtv"})
		assert.Equal(t, "YES", result)
	})
}
