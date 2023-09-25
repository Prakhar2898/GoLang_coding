package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedlist struct {
	head *Node
}

func main() {
	ll := linkedlist{}
	ll.append(3)
	ll.append(4)
	ll.append(45)
	ll.display()
}

func (l *linkedlist) append(d int) {
	node := &Node{data: d, next: nil}

	if l.head == nil {
		l.head = node
		return
	}
	current := l.head

	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *linkedlist) display() {
	current := l.head
	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println(current.data)
}
