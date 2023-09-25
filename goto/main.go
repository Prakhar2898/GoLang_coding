package main

import (
	"fmt"
	"strings"
)

func main() {
	// var i int
	// const v = 23
	// v = 23
	// fmt.Print(v)

	// Go:
	// 	fmt.Println("Here I am")

	// 	fmt.Scanln(&i)

	// 	goto Go
	v, e := squareRt(-1)
	fmt.Println(e)
	fmt.Println(v)

	fmt.Println(divideV(10, 0))
	fmt.Println(divideV(10, 10))

	s := "Heroeee"
	s1 := strings.Replace(s, "e", "t", 1)
	fmt.Print(s1)
	fileIO()

}
