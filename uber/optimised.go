package main

import "fmt"

func check() {
	n := []int{1, 2, 3, 4, 5}
	index := 0

	for i := 2; i < len(n); i++ {
		n[index], n[i] = n[i], n[index]
		if index <= 2 {
			index++
		}
	}
	fmt.Println(n)
}
