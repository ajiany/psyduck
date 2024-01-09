package tree

import "container/list"

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	runFunc := func(curNode *TreeNode, depth int) {}
	runFunc = func(curNode *TreeNode, depth int) {
		if curNode == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], curNode.Val)

		runFunc(curNode.Left, depth+1)
		runFunc(curNode.Right, depth+1)
	}

	runFunc(root, 0)

	return
}

func levelOrderV2(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmpArr := make([]int, 0)

		length := queue.Len()
		for i := 0; i < length; i++ {

			element := queue.Remove(queue.Front()).(*TreeNode)

			tmpArr = append(tmpArr, element.Val)

			if element.Left != nil {
				queue.PushBack(element.Left)
			}

			if element.Right != nil {
				queue.PushBack(element.Right)
			}

		}

		res = append(res, tmpArr)
	}

	return
}
