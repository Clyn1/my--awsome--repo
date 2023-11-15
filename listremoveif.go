package piscine

import "fmt"

type Data interface{}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func ListMerge(l1 *LinkedList, l2 *LinkedList) {
	if l1.head == nil {
		l1.head = l2.head
		return
	}

	current := l1.head
	for current.next != nil {
		current = current.next
	}

	current.next = l2.head
}

func main() {
	list1 := LinkedList{}
	list1.head = &Node{data: 1}
	list1.head.next = &Node{data: 2}

	list2 := LinkedList{}
	list2.head = &Node{data: 3}
	list2.head.next = &Node{data: 4}

	fmt.Println("Before ListMerge:")
	printList(list1)

	ListMerge(&list1, &list2)

	fmt.Println("After ListMerge:")
	printList(list1)
}

func printList(l LinkedList) {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
