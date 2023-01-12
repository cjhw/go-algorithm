package Tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	min := math.MaxInt64
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)

		if prev != nil && root.Val-prev.Val < min {
			min = root.Val - prev.Val
		}

        prev = root

		traverse(root.Right)
	}

    traverse(root)

	return min
}