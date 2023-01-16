package Tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	if p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil

}