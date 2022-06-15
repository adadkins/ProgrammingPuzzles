package super_digit

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test
// go test -fuzz=Fuzz -fuzztime 30s
func TestSuperDigit(t *testing.T) {
	t.Run("test multiple of 10 with 1 concat", func(t *testing.T) {
		result := SuperDigit("1234", 1)
		assert.Equal(t, 1, result)
	})

	t.Run("test multiple of 10 with 6 concat", func(t *testing.T) {
		result := SuperDigit("1234", 6)
		assert.Equal(t, 6, result)
	})

	t.Run("hacker ranke test case 0", func(t *testing.T) {
		result := SuperDigit("148", 3)
		assert.Equal(t, 3, result)
	})

	t.Run("hacker ranke test case 1", func(t *testing.T) {
		result := SuperDigit("9875", 4)
		assert.Equal(t, 8, result)
	})

	t.Run("hacker ranke test case 2", func(t *testing.T) {
		result := SuperDigit("123", 3)
		assert.Equal(t, 9, result)
	})
}
func FuzzSuperDigit(f *testing.F) {
	f.Add(4, 5)
	f.Fuzz(func(t *testing.T, j, i int) {
		if i < 1 || j < 0 {
			t.Skip()
		}
		strconv.Itoa(j)
		knownAlgorithmOutput := StackOverflowAlgorithm(strconv.Itoa(j), i)
		customAlgorithmOutput := SuperDigit(strconv.Itoa(j), i)
		assert.Equal(t, knownAlgorithmOutput, customAlgorithmOutput)
	})
}

// copied from stack overflow
func StackOverflowAlgorithm(n string, k int) int {
	sum := 0
	for _, x := range n {
		sum += int(x) - '0'
	}
	sum *= int(k)
	n = strconv.Itoa(sum)
	for len(n) > 1 {
		sum = 0
		for _, x := range n {
			sum += int(x) - '0'
		}
		n = strconv.Itoa(sum)
	}
	return int(sum)
}
