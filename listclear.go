package piscine


type Data interface{}package main


type Data interface{}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func ListClear(l *LinkedList) {
	l.head = nil
}

func main() {
	list := LinkedList{}
	list.head = &Node{data: 1}
	list.head.next = &Node{data: 2}
	list.head.next.next = &Node{data: 3}

	fmt.Println("Before ListClear:")
	printList(list)

	ListClear(&list)

	fmt.Println("After ListClear:")
	printList(list)
}

func printList(l LinkedList) {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func ListClear(l *LinkedList) {
	l.head = nil
}

func main() {
	list := LinkedList{}
	list.head = &Node{data: 1}
	list.head.next = &Node{data: 2}
	list.head.next.next = &Node{data: 3}

	fmt.Println("Before ListClear:")
	printList(list)

	ListClear(&list)

	fmt.Println("After ListClear:")
	printList(list)
}

func printList(l LinkedList) {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
