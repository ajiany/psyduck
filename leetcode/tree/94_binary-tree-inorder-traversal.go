package tree

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}
