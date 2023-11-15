package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func ListReverse(l *List) {
	var prev *Node
	current := l.head
	var next *Node

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

func PrintList(l *List) {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	myList := &List{
		head: &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: nil}}},
	}

	fmt.Println("Original List:")
	PrintList(myList)

	ListReverse(myList)

	fmt.Println("Reversed List:")
	PrintList(myList)
}
