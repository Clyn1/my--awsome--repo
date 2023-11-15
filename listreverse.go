package piscine

func ListLast(l *List) *NodeL {
	current := l.Head

	for current != nil && current.Next != nil {
		current = current.Next
	}

	return current
}
