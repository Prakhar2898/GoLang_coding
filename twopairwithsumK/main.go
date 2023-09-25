package main

import "fmt"

func main() {
	a := []int{1, 5, 7, -1}
	m := map[int]int{}
	sum := 6
	count := 0
	for i := 0; i < len(a); i++ {
		m[a[i]] = 1
	}
	fmt.Println(m)

	for _, v := range a {
		k := sum - v
		_, r := m[k]
		if r {
			count++
		}
	}
	fmt.Println(count / 2)

}
