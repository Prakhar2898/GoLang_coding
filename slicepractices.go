package main

import (
	"fmt"
	"sort"

	"strings"
)

func sliceCode(s []string) {
	str1 := strings.Join(s, ",")
	fmt.Printf("%T %v\n", str1, str1)
	fmt.Printf("strings.Contains(str1, \"Apple\"): %v\n", strings.Contains(str1, "Apple"))

}

func stringCode(s string) {
	str := strings.Split(s, " ")
	fmt.Println(str)
	fmt.Printf("%T\n", str)
	str = append(str, "Uber")

	fmt.Println(str)

}

func reverseSlice(n []int) []int {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func sortSlice(n []int) []int {
	sort.Ints(n)
	return n
}

func sortStringSlice(n []string) {
	sort.Strings(n)
	fmt.Println(n)
	s := n
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
