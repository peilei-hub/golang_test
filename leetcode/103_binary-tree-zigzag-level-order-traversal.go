package main

func zigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := make([]*leetcode.TreeNode, 0)
	queue = append(queue, root)

	reversed := false
	for len(queue) != 0 {
		curLevelLen := len(queue) // 当前level节点数量
		curLevelVals := make([]int, curLevelLen)
		nextLevelList := make([]*leetcode.TreeNode, 0) // 保存下一层level节点
		for i, node := range queue {
			curLevelVals[i] = node.Val
			if node.Left != nil {
				nextLevelList = append(nextLevelList, node.Left)
			}
			if node.Right != nil {
				nextLevelList = append(nextLevelList, node.Right)
			}
		}

		if reversed {
			reverse103(curLevelVals) // 翻转当前level的vals
		}
		reversed = !reversed

		result = append(result, curLevelVals)

		if len(nextLevelList) > 0 {
			queue = append(queue, nextLevelList...)
		}
		queue = queue[curLevelLen:]
	}

	return result
}

func reverse103(res []int) {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
}
