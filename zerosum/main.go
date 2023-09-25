package main

import (
	"fmt"
)

func main() {
	num := []int{15, -2, 2, -8, 1, 7, 10, 23}
	//num := []int{1, -1, 3, 2, -2, -8, 1, 7, 10, 23}
	i, j := zeroSum(num)
	num = num[i : j+1]
	fmt.Println(num)
	// s := int(math.Max(2, 3))
	// fmt.Printf("%T %d", s, s)

}

func zeroSum(n []int) (int, int) {
	mp := map[int]int{}
	start := 0
	end := 0
	sum := 0
	for i, v := range n {
		if sum < 0 {
			sum = 0
		}
		sum = sum + v

		val, ck := mp[sum]
		if ck {
			if i-mp[sum] > end-start {
				end = i
				start = val + 1
			}
		} else {
			mp[sum] = i
		}
	}
	return start, end
}
