package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	t.Run("BaseCase: Test Phone number matches one number", func(t *testing.T) {
		m := *(NewTimerMap())
		m.Set(1, 1, 1)
		result := m.Get(1, 1)
		assert.EqualValues(t, 1, result)
	})

	t.Run("Example1: set value, change value at later time, check later time has been updated", func(t *testing.T) {
		m := *(NewTimerMap())
		m.Set(1, 1, 0)
		m.Set(1, 2, 2)
		assert.EqualValues(t, 1, m.Get(1, 1))
		assert.EqualValues(t, 2, m.Get(1, 3))
	})

	t.Run("Example3: set value, change value at later time, check that time has been updated", func(t *testing.T) {
		m := *(NewTimerMap())
		m.Set(1, 1, 0)
		m.Set(1, 2, 0)
		assert.EqualValues(t, 2, m.Get(1, 0))
	})
}
