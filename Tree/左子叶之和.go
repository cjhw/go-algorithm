package Tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := sumOfLeftLeaves(root.Left)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left = root.Left.Val
	}

	right := sumOfLeftLeaves((root.Right))

	return left + right
}
