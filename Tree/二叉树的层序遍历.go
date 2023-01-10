package Tree

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	if root == nil {
		return res
	}
	for len(queue) > 0 {
		length := len(queue)
		curLevel := []int{}
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if (node.Right) != nil {
				queue = append(queue, node.Right)
			}
			curLevel = append(curLevel, node.Val)
		}
		res = append(res, curLevel)
		curLevel = []int{}
	}

	return res
}
