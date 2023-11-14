package piscine

type Data interface{}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func ListLast(l LinkedList) Data {
	if l.head == nil {
		return nil
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	return current.data
}
