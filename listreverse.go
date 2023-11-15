package piscine

type Node struct {
	data int
	next *Node
}

type List struct {
	Head *Node
}

func ListReverse(l *List) {
	var prev *Node
	current := l.Head
	var next *Node

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
