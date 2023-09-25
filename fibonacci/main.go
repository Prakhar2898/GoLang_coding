package main

import "fmt"

func main() {
	n := 8
	fibS := []int{0, 1}

	for i := 3; i <= n; i++ {
		l := len(fibS)
		fibS = append(fibS, fibS[l-1]+fibS[l-2])
	}
	fmt.Println(fibS)

}
