package tree

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}

	res = append(res, root.Val)

	return res
}
