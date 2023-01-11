package Tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := sumOfLeftLeaves(root.Left)

	right := sumOfLeftLeaves((root.Right))

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left = root.Left.Val
	}

	return left + right
}
