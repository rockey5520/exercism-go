package isogram

import (
	"strings"
)

// IsIsogram check given string if it is isogram or not
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	for _, v := range input {

		if strings.Count(input, string(v)) > 1 && v >= 'a' && v <= 'z' {
			return false
		}
	}
	return true
}