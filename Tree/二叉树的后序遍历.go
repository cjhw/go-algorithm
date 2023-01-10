package Tree

import "container/list"

// 递归
// func postorderTraversal(root *TreeNode) (res []int) {
// 	var traversal func(node *TreeNode)
// 	traversal = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		traversal(node.Left)
// 		traversal(node.Right)
// 		res = append(res, node.Val)
// 	}
// 	traversal(root)
// 	return res
// }

// 迭代
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	reverse(ans)
	return ans
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
