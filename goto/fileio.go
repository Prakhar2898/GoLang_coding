package main

import (
	"fmt"
	"io/ioutil"
)

func fileIO() {
	s := "Hi, there"
	err := ioutil.WriteFile("myfile", []byte(s), 0666)
	if err != nil {
		fmt.Println(err)
	}
	b, e := ioutil.ReadFile("myfile")
	if e != nil {
		fmt.Println(e)
	}
	b1, _ := ioutil.ReadFile("myfile")

	fmt.Println(string(b))

	fmt.Println(string(b1))

}
