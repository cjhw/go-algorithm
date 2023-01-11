package Tree

func findMax(nums []int) (index int) {
	for i, v := range nums {
		if nums[index] < v {
			index = i
		}
	}
	return
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	index := findMax(nums)

	root := &TreeNode{
		Val:   nums[index],
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}

	return root
}
