package programmingpuzzles_test

import (
	programmingpuzzles "programmingPuzzles"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCeasarCipher(t *testing.T) {
	t.Run("test cipher", func(t *testing.T) {
		shifted := programmingpuzzles.CeasarCipher("abcde", 1)
		assert.Equal(t, "bcdef", shifted)
	})
	t.Run("test cipher capital letters", func(t *testing.T) {
		shifted := programmingpuzzles.CeasarCipher("ABCDE", 1)
		assert.Equal(t, "BCDEF", shifted)
	})
	t.Run("test cipher capital and lowercase letters", func(t *testing.T) {
		shifted := programmingpuzzles.CeasarCipher("AbCdEfGhIjKlMnOpQrStUvWxYz", 1)
		assert.Equal(t, "BcDeFgHiJkLmNoPqRsTuVwXyZa", shifted)
	})

	t.Run("test cipher capital and lowercase letters with a large shift", func(t *testing.T) {
		shifted := programmingpuzzles.CeasarCipher("AbCdEfGhIjKlMnOpQrStUvWxYz", 25)
		assert.Equal(t, "ZaBcDeFgHiJkLmNoPqRsTuVwXy", shifted)
	})
	t.Run("test cipher capital and lowercase letters with a large shift", func(t *testing.T) {
		shifted := programmingpuzzles.CeasarCipher("ABCDEFGHIjklmnopqrstuvwxyZ", 25)
		assert.Equal(t, "ZaBcDeFgHiJkLmNoPqRsTuVwXy", shifted)
	})
	// t.Run("test cipher with real sentence", func(t *testing.T) {
	// 	shifted := programmingpuzzles.CeasarCipher("The-red-cat-Fell-Out-", 1)
	// 	assert.Equal(t, "BcDeFgHiJkLmNoPqRsTuVwXyZa", shifted)
	// })
}
