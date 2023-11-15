package piscine

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
