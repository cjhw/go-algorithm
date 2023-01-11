package Tree

var hash map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	hash = make(map[int]int)

	for i, v := range inorder {
		hash[v] = i
	}

	return rebuild(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
}

func rebuild(inorder []int, postorder []int, rootIdx int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r { // 只剩唯一一个元素，直接返回
		return &TreeNode{Val: inorder[l]}
	}
	rootV := postorder[rootIdx]
	rootIn := hash[rootV]
	root := &TreeNode{Val: inorder[rootIn]}

	root.Left = rebuild(inorder, postorder, rootIdx-(r-rootIn)-1, l, rootIn-1)
	root.Right = rebuild(inorder, postorder, rootIdx-1, rootIn+1, r)

	return root
}
