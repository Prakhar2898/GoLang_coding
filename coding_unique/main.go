package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s1 := []string{"a", "a", "b", "a", "c", "d", "b", "c", "e"}
	mp := map[string]int{}
	sort.Strings(s1)
	for _, v := range s1 {
		val, b := mp[v]
		if b {
			mp[v] = val + 1
		} else {
			mp[v] = 1
		}
	}
	fmt.Println(mp)

	s2 := ""

	for k, v := range mp {
		s2 = s2 + k + strconv.Itoa(v)
	}

	fmt.Println(s2)

}
