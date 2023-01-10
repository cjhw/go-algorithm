package Tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// func preorderTraversal(root *TreeNode) (res []int) {
// 	var traversal func(node *TreeNode)
// 	traversal = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		res = append(res, node.Val)
// 		traversal(node.Left)
// 		traversal(node.Right)
// 	}
// 	traversal(root)
// 	return res
// }

// 迭代
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		res = append(res, node.Val)

		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}

	}

	return res
}
