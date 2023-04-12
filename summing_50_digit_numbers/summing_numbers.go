package sum_numbers

import (
	"fmt"
	"strconv"
)

func SumNumbers(input []string) string {
	bigNumber := newBigNumber()

	for _, v := range input {
		bigNumber.add(v)
	}

	return bigNumber.getFirstTenDigitsOfSum()
}

type BigNumber struct {
	sum []int
}

func newBigNumber() *BigNumber {
	arr := []int{}
	for i := 0; i < 50; i++ {
		arr = append(arr, 0)
	}

	return &BigNumber{
		sum: arr,
	}
}

func (n *BigNumber) add(input string) {
	carriedOnesPlace := 0
	for i := 0; i < len(n.sum); i++ {
		v := len(input) - i - 1
		inputDigit := 0
		if v >= 0 {
			inputDigit, _ = strconv.Atoi(input[v : v+1])
		}

		newDigit := inputDigit + n.sum[i] + carriedOnesPlace
		carriedOnesPlace = 0
		n.sum[i] = newDigit % 10
		if newDigit/10 >= 1 {
			carriedOnesPlace = 1
		}
	}

	if carriedOnesPlace > 0 {
		n.sum = append(n.sum, carriedOnesPlace%10)
		carriedOnesPlace = carriedOnesPlace / 10
	}
}

func (n *BigNumber) getFirstTenDigitsOfSum() string {
	returnString := ""
	for i := len(n.sum) - 1; i > len(n.sum)-10-1 && i >= 0; i-- {
		returnString = fmt.Sprintf("%v%v", returnString, strconv.Itoa(n.sum[i]))
	}
	return returnString
}
