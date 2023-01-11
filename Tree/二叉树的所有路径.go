package Tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var traverse func(root *TreeNode, s string)

	traverse = func(root *TreeNode, s string) {
		if root.Left == nil && root.Right == nil {
			s = s + strconv.Itoa(root.Val)
			res = append(res, s)
			return
		}

		s = s + strconv.Itoa(root.Val) + "->"

		if root.Left != nil {
			traverse(root.Left, s)
		}

		if root.Right != nil {
			traverse(root.Right, s)
		}
	}

	traverse(root, "")
	return res
}
