package piscine

import "os"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func ApplyPreorder(root *TreeNode, f func(int)) {
	if root != nil {
		f(root.Value)
		ApplyPreorder(root.Left, f)
		ApplyPreorder(root.Right, f)
	}
}

func printNode(val int) {
	str := string(rune(val))
	os.Stdout.WriteString(str)
}

func main() {
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 3,
			},
			Right: &TreeNode{
				Value: 4,
			},
		},
		Right: &TreeNode{
			Value: 5,
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}

	ApplyPreorder(root, func(val int) {
		printNode(val)
	})
}
