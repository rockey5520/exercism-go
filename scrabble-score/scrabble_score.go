package scrabble

import (
	"strings"
)

// Score function provides scrabble score
func Score(s string) int {
	scrabbleMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		var scrabbleCharacter = strings.ToUpper(string(s[i]))
		var count = strings.Count(strings.ToUpper(s), scrabbleCharacter)
		_, present := scrabbleMap[scrabbleCharacter]
		if !present {
			scrabbleMap[scrabbleCharacter] = count
		}
	}
	var score int
	masterScore := map[string]int{"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1, "D": 2, "G": 2, "B": 3, "C": 3, "M": 3, "P": 3, "F": 4, "H": 4, "V": 4, "W": 4, "Y": 4, "K": 5, "J": 8, "X": 8, "Q": 10, "Z": 10}
	for k, v := range scrabbleMap {
		_, present := masterScore[k]
		if present {
			var charScore = masterScore[k]
			score = score + (v * charScore)
		}
	}

	return score

}
