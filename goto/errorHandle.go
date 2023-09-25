package main

import (
	"errors"
	"math"
)

func squareRt(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Value can not be zero")
	}
	return int(math.Sqrt(float64(n))), nil
}
