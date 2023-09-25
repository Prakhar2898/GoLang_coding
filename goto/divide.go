package main

import "fmt"

func divideV(n1, n2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	q := n1 / n2
	return q
}
