package Tree

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var traverse func(node1, node2 *TreeNode) bool

	traverse = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}

		isSym1 := traverse(node1.Left, node2.Right)
		isSym2 := traverse(node1.Right, node2.Left)

		return isSym1 && isSym2
	}

	return traverse(root.Left, root.Right)
}
