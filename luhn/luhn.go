package luhn

import (
	"strconv"
	s "strings"
)

// Valid function validates given string per luhn formula
func Valid(input string) bool {
	input = s.Replace(input, " ", "", -1)
	if len(input) < 2 {
		return false
	}
	var digit, doubeSum, sum int
	var err error
	for i := 0; i < len(input); i++ {
		digit, err = strconv.Atoi(string(input[len(input)-1-i]))
		if err != nil {
			return false
		}
		if i%2 == 0 {
			sum += digit
			continue
		}
		doubeSum = 2 * digit
		if doubeSum > 9 {
			doubeSum -= 9
		}
		sum += doubeSum
	}
	return sum%10 == 0
}
