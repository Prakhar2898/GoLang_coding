package main

import (
	"fmt"
	"os"
)

func main() {
	n := []int{1, 2, 3, 4, 5}
	// 3 4 5 1 2
	//2
	var k int
	fmt.Scanln(&k)
	k = k % len(n)
	if k > len(n) {
		fmt.Println("Input value is greater than length of array")
		os.Exit(1)
	}

	n1 := []int{}
	sum := 0
	for i := k; i < len(n); i++ {
		n1 = append(n1, n[i])
		if i%2 == 0 {
			sum = sum + n[i]
		}
	}
	for i := 0; i < k; i++ {
		n1 = append(n1, n[i])
		if i%2 == 0 {
			sum = sum + n[i]
		}

	}
	fmt.Println(n1)
	fmt.Println(sum)
	check()

}
