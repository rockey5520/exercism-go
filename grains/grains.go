package grains

import (
	"errors"
)

//Square returns number of grains on a particular square
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("can't calculate value outside of square, must be between 1 and 64")
	}
	return 1 << (i - 1), nil
}

//Total returns number of grains on the board
func Total() uint64 {
	return 1<<64 - 1
}
