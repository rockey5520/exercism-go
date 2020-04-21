package raindrops

import (
	"strconv"
)

// Convert function converts integer to rain sounds
func Convert(a int) string {
	var s string
	if a%3 == 0 {
		s += "Pling"
	}
	if a%5 == 0 {
		s += "Plang"
	}
	if a%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s += strconv.Itoa(a)
		return s
	}
	return s
}