package Tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{
			Val: val,
		}

		return node
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}