// Package leap had method to calculate leap year.
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(a int) bool {
	if a%4 == 0 {
		if a%100 == 0 {
        return  a%400 == 0
		}
		return true
	}
	return false
}
