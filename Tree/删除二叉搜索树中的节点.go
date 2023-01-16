package Tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	minNode := root.Right

	for minNode.Left != nil {
		minNode = minNode.Left
	}

	root.Right = deleteNode1(root.Right)

	root.Val = minNode.Val
	return root
}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteNode1(root.Left)
	return root
}