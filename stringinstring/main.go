package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "I lov Love my India"
	s2 := " lov"
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	result := containsString(s1, s2)
	fmt.Println(result)

}

func containsString(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[0] {
			s3 := s1[i : len(s2)+i]
			if s3 == s2 {
				return true
			}
		}
	}
	return false
}
