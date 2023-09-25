package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxi := math.MinInt
	sum := 0
	for _, v := range arr {
		sum += v
		if sum > maxi {
			maxi = sum
		}
		if sum < 0 {
			sum = 0
		}

	}
	fmt.Println(maxi)
}
