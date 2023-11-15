package piscine

type NodeL struct {
	Data int
	Next *NodeL
}

type List struct {
	Head *NodeL
}

func ListReverse(l *List) {
	var prev, next *NodeL
	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
