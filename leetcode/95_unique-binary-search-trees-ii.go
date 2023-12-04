package main

// https://leetcode.cn/problems/unique-binary-search-trees-ii/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	trees := allTrees(1, n)

	return trees
}

// 获得 [start,end]的树列表
func allTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // 数组里有一个空树
	}
	if start == end { // 可以不加，下方的遍历会包含
		return []*TreeNode{&TreeNode{Val: start}}
	}

	result := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		leftTrees := allTrees(start, i-1)
		rightTrees := allTrees(i+1, end)

		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				curTree := &TreeNode{
					Val:   i,
					Left:  leftTree,
					Right: rightTree,
				}

				result = append(result, curTree)
			}
		}
	}

	return result
}
