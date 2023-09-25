package main

import (
	"fmt"
)

func main() {
	num := []int{12, 23, 34, 45, 56, 67, 78, 89, 91, 92}

	fmt.Println(num)

	i := 3
	num = append(num[:i], num[i+1:]...)

	fmt.Println(num)

}
