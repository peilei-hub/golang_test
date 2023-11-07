package main

func minDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0

	queue := make([]*leetcode.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		depth++
		curLen := len(queue)
		nextLevelNodes := make([]*leetcode.TreeNode, 0)
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}

			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		queue = append(queue, nextLevelNodes...)
		queue = queue[curLen:]
	}

	return depth
}
