package Tree

import "strconv"

// func binaryTreePaths(root *TreeNode) []string {
// 	res := make([]string, 0)

// 	var traverse func(root *TreeNode, s string)

// 	traverse = func(root *TreeNode, s string) {
// 		if root.Left == nil && root.Right == nil {
// 			s = s + strconv.Itoa(root.Val)
// 			res = append(res, s)
// 			return
// 		}

// 		s = s + strconv.Itoa(root.Val) + "->"

// 		if root.Left != nil {
// 			traverse(root.Left, s)
// 		}

// 		if root.Right != nil {
// 			traverse(root.Right, s)
// 		}
// 	}

// 	traverse(root, "")
// 	return res
// }

func binaryTreePaths(root *TreeNode) []string {
	stack := []*TreeNode{}
	paths := make([]string, 0)
	res := make([]string, 0)
	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}

	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		path := paths[l-1]
		stack = stack[:l-1]
		paths = paths[:l-1]

		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}

	return res
}
