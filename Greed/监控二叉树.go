package Greed

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {

	// 	0：该节点无覆盖
	// 1：本节点有摄像头
	// 2：本节点有覆盖

	res := 0

	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 2
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		// 情况1
		// 左右节点都有覆盖
		if left == 2 && right == 2 {
			return 0
		}

		// 情况2
		// left == 0 && right == 0 左右节点无覆盖
		// left == 1 && right == 0 左节点有摄像头，右节点无覆盖
		// left == 0 && right == 1 左节点有无覆盖，右节点摄像头
		// left == 0 && right == 2 左节点无覆盖，右节点覆盖
		// left == 2 && right == 0 左节点覆盖，右节点无覆盖
		if left == 0 || right == 0 {
			res++
			return 1
		}

		// 情况3
		// left == 1 && right == 2 左节点有摄像头，右节点有覆盖
		// left == 2 && right == 1 左节点有覆盖，右节点有摄像头
		// left == 1 && right == 1 左右节点都有摄像头
		if left == 1 || right == 1 {
			return 2
		}

		// 不会走到这里
		return -1
	}

	// 情况4
	if dfs(root) == 0 {
		// root 无覆盖
		res++
	}

	return res
}
