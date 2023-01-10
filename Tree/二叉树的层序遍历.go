package Tree

// 递归
func levelOrder(root *TreeNode) [][]int {
	var traversel func(root *TreeNode, depth int)
	var res [][]int
	depth := 0

	traversel = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		traversel(root.Left, depth+1)
		traversel(root.Right, depth+1)
	}

	traversel(root, depth)

	return res
}

// 队列
// func levelOrder(root *TreeNode) [][]int {
// 	var res [][]int
// 	queue := []*TreeNode{root}
// 	if root == nil {
// 		return res
// 	}
// 	for len(queue) > 0 {
// 		length := len(queue)
// 		curLevel := []int{}
// 		for i := 0; i < length; i++ {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if (node.Right) != nil {
// 				queue = append(queue, node.Right)
// 			}
// 			curLevel = append(curLevel, node.Val)
// 		}
// 		res = append(res, curLevel)
// 		curLevel = []int{}
// 	}

// 	return res
// }
