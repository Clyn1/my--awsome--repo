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
		// Type switch to check the type of data
		switch data := l2.Head.Data.(type) {
		case int:
			ListPushBack(l1, data)
		case []interface{}:
			// Handle slices by appending their elements
			for _, item := range data {
				ListPushBack(l1, item)
			}
		default:
			// Handle other types accordingly
			ListPushBack(l1, data)
		}

		l2.Head = l2.Head.Next
	}
}
