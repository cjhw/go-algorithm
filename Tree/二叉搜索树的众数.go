package Tree

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode

	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {

		if root == nil {
			return
		}

		traverse(root.Left)

		if prev != nil && prev.Val == root.Val {
			count++
		} else {
			count = 1
		}

		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{root.Val}
			} else {
				res = append(res, root.Val)
			}
			max = count
		}

		prev = root

		traverse(root.Right)
	}

	traverse(root)

	return res
}