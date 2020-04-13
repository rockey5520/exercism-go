package hamming

import (
	"errors"
	s "strings"
)

// Distance function calculates the hamming distance between two string
func Distance(a, b string) (int, error) {
	hammingdistance := 0
	if len(a) != len(b) {
		return 0, errors.New("strands are of different length")
	}
	if s.Compare(s.ToLower(a), s.ToLower(b)) == 0 {
		return 0, nil
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingdistance++
		}
	}
	return hammingdistance, nil
}
