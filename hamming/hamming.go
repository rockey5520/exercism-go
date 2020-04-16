package hamming

import (
	"errors"
)

// Distance function calculates the hamming distance between two string
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("strands are of different length")
	}
	runeA := []rune(a)
	runeB := []rune(b)
	var hammingdistance int
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			hammingdistance++
		}
	}
	return hammingdistance, nil
}
