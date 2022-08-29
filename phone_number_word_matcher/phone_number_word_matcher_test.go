package programmingpuzzles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCeasarCipher(t *testing.T) {
	t.Run("Test Phone number matches one number", func(t *testing.T) {
		phoneNumbers := []int{1234567890, 623277537, 4807767870}
		result := PhoneWordMatcher("apples", phoneNumbers)
		assert.EqualValues(t, []int{623277537}, result, "Word should have been in the list of phone numbers")
		assert.Len(t, result, 1, "One phone number should have been found")
		assert.EqualValues(t, result[0], 623277537, "Phone number found should have been this number")
	})
	t.Run("Test Phone number matches no numbers", func(t *testing.T) {
		phoneNumbers := []int{1110001100, 0001110011, 0000000000, 1111111111}
		result := PhoneWordMatcher("apples", phoneNumbers)
		assert.EqualValues(t, []int{}, result, "Word should have NOT been in the list of phone numbers")
		assert.Len(t, result, 0, "One phone number should have been found")
	})

	t.Run("Test Phone number matches more than one number", func(t *testing.T) {
		phoneNumbers := []int{1234567890, 623277537, 4807767870, 277537123}
		result := PhoneWordMatcher("apples", phoneNumbers)
		assert.EqualValues(t, []int{623277537, 277537123}, result, "Word should have been in the list of phone numbers")
		assert.Len(t, result, 2, "Two phone numbers should have been found")
		assert.EqualValues(t, result[0], 623277537, "Phone number found should have been this number")
		assert.EqualValues(t, result[1], 277537123, "Phone number found should have been this number")

	})

	t.Run("Large string matches ie 10 digits long", func(t *testing.T) {
		phoneNumbers := []int{7872923779, 623277537, 4807767870}
		result := PhoneWordMatcher("strawberry", phoneNumbers)
		assert.EqualValues(t, []int{7872923779}, result, "Word should have been in the list of phone numbers")
		assert.Len(t, result, 1, "One phone number should have been found")
		assert.EqualValues(t, result[0], 7872923779, "Phone number found should have been this number")
	})

	t.Run("Small string matches ie 1 digit long", func(t *testing.T) {
		phoneNumbers := []int{1234567890, 623277537, 4807767870}
		result := PhoneWordMatcher("a", phoneNumbers)
		assert.EqualValues(t, []int{1234567890, 623277537}, result, "Word should have been in the list of phone numbers")
		assert.Len(t, result, 2, "Two phone numbers should have been found")
		assert.EqualValues(t, result[0], 1234567890, "Phone number found should have been this number")
		assert.EqualValues(t, result[1], 623277537, "Phone number found should have been this number")
	})

	t.Run("UPPER case words are translated", func(t *testing.T) {
		phoneNumbers := []int{7872923779, 623277537, 4807767870}
		result := PhoneWordMatcher("STRAWBERRY", phoneNumbers)
		assert.EqualValues(t, []int{7872923779}, result, "Word should have been in the list of phone numbers")
		assert.Len(t, result, 1, "One phone number should have been found")
		assert.EqualValues(t, result[0], 7872923779, "Phone number found should have been this number")
	})

	t.Run("Parse word into phone digits test", func(t *testing.T) {
		result := convertWordToPhoneNumber("aaa")
		assert.EqualValues(t, "222", result)
	})
	t.Run("Parse word into phone digits test with zero values", func(t *testing.T) {
		result := convertWordToPhoneNumber(" .!%.")
		assert.EqualValues(t, "", result)
	})
}
