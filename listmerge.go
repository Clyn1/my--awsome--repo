package piscine

func ListMerge(l1 *LinkedList, l2 *LinkedList) {
	for l2.head != nil {
		ListPushBack(l1, l2.head.Data)
		l2.head = l2.head.Next
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
