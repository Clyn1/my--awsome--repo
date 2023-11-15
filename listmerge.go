package piscine

// ListPushBack inserts a new element at the end of the list
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// ListMerge merges two lists by appending elements from l2 to l1
func ListMerge(l1 *List, l2 *List) {
	for l2.Head != nil {
		// Type assertion to convert interface{} to int
		data, ok := l2.Head.Data.(int)
		if !ok {
			// Handle the case where the assertion fails
			// For now, we'll assume the data is always an int
			panic("Data is not of type int")
		}

		ListPushBack(l1, data)
		l2.Head = l2.Head.Next
	}
}
