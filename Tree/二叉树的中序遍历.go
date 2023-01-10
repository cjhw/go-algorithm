package Tree

import "container/list"

// 递归
// func inorderTraversal(root *TreeNode) (res []int) {
// 	var traversal func(node *TreeNode)
// 	traversal = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		traversal(node.Left)
// 		res = append(res, node.Val)
// 		traversal(node.Right)
// 	}

// 	traversal(root)
// 	return res
// }

// 迭代
func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	cur := root
	st := list.New()

	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}

	return res
}
