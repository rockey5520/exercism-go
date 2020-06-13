package isogram

import (
	"strings"
	s "strings"
)

// IsIsogram check given string if it is isogram or not
func IsIsogram(input string) bool {
	input = s.Replace(input, "-", "", -1)
	input = s.Replace(input, " ", "", -1)
	input = strings.ToLower(input)

	for _, v := range input {
		if strings.Count(input, string(v)) > 1 {
			return false
		}
	}
	return true
}
