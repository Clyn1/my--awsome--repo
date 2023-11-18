package piscine

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

func printNode(val int, buffer *[]byte) {
	*buffer = append(*buffer, byte(val+'0'))
}

func printBuffer(buffer []byte) {
	for _, char := range buffer {
		print(char)
	}
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

	var outputBuffer []byte

	ApplyPreorder(root, func(val int) {
		printNode(val, &outputBuffer)
	})

	printBuffer(outputBuffer)
}
