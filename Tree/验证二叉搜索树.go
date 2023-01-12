package Tree

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var traverse func(root *TreeNode) bool
	traverse = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		left := traverse(root.Left)

		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev = root

		right := traverse(root.Right)

		return left && right
	}

	return traverse(root)
}