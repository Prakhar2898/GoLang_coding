package main

import "fmt"

func main() {
	fmt.Println(palindrome(121))
	stringP()
}
func palindrome(n int) bool {
	x := n
	j := 0

	for x != 0 {
		r := x % 10
		x = x / 10
		j = j*10 + r
	}
	if n == j {
		return true
	}
	return false
}
