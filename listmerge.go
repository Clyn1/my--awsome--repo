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
	for l2.head != nil {
		ListPushBack(l1, l2.head.data)
		l2.head = l2.head.next
	}
}

func ListPushBack(l *LinkedList, data Data) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}
