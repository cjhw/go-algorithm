package Dp

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const inf = math.MaxInt32 / 2

func minCameraCover(root *TreeNode) int {

	// a： root必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
	// b：覆盖整棵树需要的摄像头数目，无论无论 root是否放置摄像头。
	// c：覆盖两棵子树需要的摄像头数目，无论节点 root本身是否被监控到。

	var dfs func(*TreeNode) (a, b, c int)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dfs = func(node *TreeNode) (a int, b int, c int) {
		if node == nil {
			return inf, 0, 0
		}

		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = 1 + lc + rc
		b = min(a, min(la+rb, ra+lb))
		c = min(a, lb+rb)
		return
	}

	_, res, _ := dfs(root)

	return res
}
