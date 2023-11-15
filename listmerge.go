package piscine

import "fmt"

type Data interface{}

type NodeL struct {
	Data Data
	Next *NodeL
}

type LinkedList struct {
	Head *NodeL
}

func ListAt(l *LinkedList, pos int) *NodeL {
	if l == nil || l.Head == nil || pos < 0 {
		return nil
	}

	current := l.Head
	for i := 0; i < pos; i++ {
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}

	return current
}

func main() {
	list := LinkedList{}
	list.Head = &NodeL{Data: 1}
	list.Head.Next = &NodeL{Data: 2}
	list.Head.Next.Next = &NodeL{Data: 3}

	pos1Node := ListAt(&list, 1)

	if pos1Node != nil {
		fmt.Println("Node at position 1:", pos1Node.Data)
	} else {
		fmt.Println("Position 1 is out of bounds.")
	}

	pos5Node := ListAt(&list, 5)

	if pos5Node != nil {
		fmt.Println("Node at position 5:", pos5Node.Data)
	} else {
		fmt.Println("Position 5 is out of bounds.")
	}
}
