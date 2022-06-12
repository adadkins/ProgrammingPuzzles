package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test
// go test -fuzz=Fuzz -fuzztime 30s
func TestCeasarCipher(t *testing.T) {
	t.Run("test cipher", func(t *testing.T) {
		shifted := CeasarCipher("abcde", 1)
		assert.Equal(t, "bcdef", shifted)
	})
	t.Run("test cipher capital letters", func(t *testing.T) {
		shifted := CeasarCipher("ABCDE", 1)
		assert.Equal(t, "BCDEF", shifted)
	})
	t.Run("test cipher capital and lowercase letters", func(t *testing.T) {
		shifted := CeasarCipher("AbCdEfGhIjKlMnOpQrStUvWxYz", 1)
		assert.Equal(t, "BcDeFgHiJkLmNoPqRsTuVwXyZa", shifted)
	})

	t.Run("test cipher capital and lowercase letters with a large shift", func(t *testing.T) {
		shifted := CeasarCipher("AbCdEfGhIjKlMnOpQrStUvWxYz", 25)
		assert.Equal(t, "ZaBcDeFgHiJkLmNoPqRsTuVwXy", shifted)
	})
	t.Run("test cipher capital and lowercase letters with a large shift", func(t *testing.T) {
		shifted := CeasarCipher("ABCDEFGHIjklmnopqrstuvwxyZ", 25)
		assert.Equal(t, "ZABCDEFGHijklmnopqrstuvwxY", shifted)
	})
	t.Run("test cipher with real sentence", func(t *testing.T) {
		shifted := CeasarCipher("The-red-cat-Fell-Out-Of-The-Tree", 1)
		assert.Equal(t, "Uif-sfe-dbu-Gfmm-Pvu-Pg-Uif-Usff", shifted)
	})
}

// this test
func FuzzCeasarCipher(f *testing.F) {
	f.Add("testtesttest", 5)
	f.Fuzz(func(t *testing.T, s string, i int) {
		if i < 0 {
			t.Skip()
		}
		knownAlgorithmOutput := StackOverflowAlgorithm(s, int32(i))
		customAlgorithmOutput := CeasarCipher(s, i)
		assert.Equal(t, knownAlgorithmOutput, customAlgorithmOutput)
	})
}

// copied from stackoverflow
func StackOverflowAlgorithm(s string, k int32) string {
	runes := []rune(s)
	if k > 26 {
		k = k % 26
	}

	for pos, char := range runes {
		if char >= 65 && char <= 90 {
			newChar := char + k
			if newChar > 90 {
				newChar = 65 + newChar - 90 - 1
			}
			runes[pos] = newChar
		}
		if char >= 97 && char <= 122 {
			newChar := char + k
			if newChar > 122 {
				newChar = 97 + newChar - 122 - 1
			}
			runes[pos] = newChar
		}
	}

	return string(runes)
}
