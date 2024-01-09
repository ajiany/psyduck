package tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	runFunc := func(curNode *TreeNode) {}
	runFunc = func(curNode *TreeNode) {
		if curNode == nil {
			return
		}

		res = append(res, curNode.Val)
		runFunc(curNode.Left)
		runFunc(curNode.Right)
	}
	runFunc(root)
	return
}
