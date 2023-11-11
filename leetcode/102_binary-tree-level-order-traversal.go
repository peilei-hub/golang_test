package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		curLevelLen := len(queue)
		nextLevelList := make([]*TreeNode, 0)
		curLevelVals := make([]int, 0)

		for _, node := range queue {
			curLevelVals = append(curLevelVals, node.Val)
			if node.Left != nil {
				nextLevelList = append(nextLevelList, node.Left)
			}
			if node.Right != nil {
				nextLevelList = append(nextLevelList, node.Right)
			}
		}

		result = append(result, curLevelVals)

		queue = append(queue, nextLevelList...)
		queue = queue[curLevelLen:]
	}

	return result
}
