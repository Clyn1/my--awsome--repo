package piscine

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
