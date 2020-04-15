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
	hammingdistance := 0
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != runeB[i] {
			hammingdistance++
		} else {
			continue
		}
	}
	if hammingdistance == 0 {
		return 0, nil
	}
	return hammingdistance, nil
}
