package Tree

func isBalanced(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := getHeight(root.Left), getHeight(root.Right)

	if l == -1 || r == -1 {
		return -1
	}

	if l-r > 1 || r-l > 1 {
		return -1
	}

	return max(l, r) + 1
}
