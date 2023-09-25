package main

import "fmt"

func swap() {
	a := 32
	b := 43
	a, b = b, a
	fmt.Println(a, " ", b)
}
