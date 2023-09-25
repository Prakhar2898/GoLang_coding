package main

import "fmt"

type engBot struct{}
type spBot struct{}

type greeting interface {
	greetingUser() string
}

func main() {
	eb := engBot{}
	sp := spBot{}

	fmt.Println(getGreet(eb))
	getGreet(sp)
	emptyInt()
}

func (e engBot) greetingUser() string {
	return "Hello"
}

func (s spBot) greetingUser() string {
	return "Halo!"
}

func getGreet(g greeting) string {
	return g.greetingUser()
}
