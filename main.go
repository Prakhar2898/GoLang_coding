package main

import (
	"fmt"
)

func main() {
	// var s string
	// s = "Hello"
	// fmt.Println(s)
	// fmt.Println(len(s))

	// result := ""
	// for i := 0; i < len(s); i++ {
	// 	result = string(s[i]) + result
	// }
	// fmt.Print(result)
	// m := map[string]int{}
	// m["Prakhar"] = 2
	// k, v := m["Prakha"]
	// fmt.Print(k, v)
	// fmt.Println(m["Prakha"])

	fmt.Println("Hello")
	// countOccurance("Hello")

	// isPalindrome := checkPalindrome("roor")
	// if isPalindrome {
	// 	fmt.Print("Yes")
	// } else {
	// 	fmt.Print("No")
	// }

	// st := "I Love India"

	// stringCode(st)

	// sl := []string{"Apple", "Banana", "Oranges", "Anaar"}
	// // sliceCode(sl)

	// intArr := []int{21, 32, 43, 54, 65, 76, 87, 98}
	// reversedArr := reverseSlice(intArr)
	// fmt.Println(reversedArr)
	// fmt.Println(sortSlice(intArr))

	// sortStringSlice(sl)

	// v1, v2 := twoSum(intArr, 130)
	// fmt.Println(v1, " ", v2)

	// swap()
	//symbol := []string{"[", "(", ")", "]", "{", "}", "{", "[", "(", ")", "(", ")", "]", "(", ")", "}"}
	symbol1 := []string{"[", "(", "]", ")"}
	result := checkPr(symbol1)
	fmt.Println(result)

}
