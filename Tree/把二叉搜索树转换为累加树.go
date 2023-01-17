package Tree

func convertBST(root *TreeNode) *TreeNode {
	pre := 0

	var traverse func(cur *TreeNode)

	traverse = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traverse(cur.Right)
		cur.Val += pre
		pre = cur.Val
		traverse(cur.Left)
	}

	traverse(root)

	return root
}
