package piscine

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
