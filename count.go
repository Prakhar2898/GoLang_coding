package main

import "fmt"

func countOccurance(s string) {
	m := map[string]int{}
	for _, v := range s {
		val, r := m[string(v)]
		if r {
			m[string(v)] = val + 1
		} else {
			m[string(v)] = 1
		}
	}
	fmt.Println(m)
}
