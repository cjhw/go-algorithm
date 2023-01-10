package Tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 1

	res += countNodes(root.Left)
	res += countNodes(root.Right)

	return res
}
