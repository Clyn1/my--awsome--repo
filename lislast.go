package piscine

type Data interface{}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) ListLast() Data {
	current := ll.head

	if current == nil {
		return nil
	}

	for current.next != nil {
		current = current.next
	}

	return current.data
}
