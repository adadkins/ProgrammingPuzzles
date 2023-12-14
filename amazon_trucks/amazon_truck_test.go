package amazon_trucks_test

import (
	"programmingpuzzles/amazon_trucks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmazonTrucks(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		items := []int{1, 3, 5, 2, 3, 2}
		trucks := []int{10, 1, 2, 3}
		result := amazon_trucks.GetTrucksForItems(trucks, items)
		expectedResult := []int{2, 0, 0, 3, 0, 3}
		assert.Equal(t, expectedResult, result)
	})
}
