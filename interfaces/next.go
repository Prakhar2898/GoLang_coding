package main

import "fmt"

func emptyInt() {
	i := []int{21, 32, 43}
	s := []string{"ps", "as", "chita"}
	printI(i)
	printI(s)

}

func printI(i interface{}) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
}
