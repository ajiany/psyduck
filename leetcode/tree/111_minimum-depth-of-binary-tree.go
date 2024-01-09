package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHigh := minDepth(root.Left)
	rightHigh := minDepth(root.Right)

	if root.Left != nil && root.Right == nil {
		return leftHigh + 1
	}

	if root.Left == nil && root.Right != nil {
		return rightHigh + 1
	}

	return min(leftHigh, rightHigh) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
