package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	n := len(preorder)

	indexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		indexMap[v] = i
	}

	return myBuildTree(preorder, inorder, indexMap, 0, n-1, 0, n-1)
}

func myBuildTree(preorder []int, inorder []int, indexMap map[int]int, preorderLeft, preorderRight, inorderLeft, inorderRight int) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}

	// 前序遍历的第一个节点就是根节点
	preOrderRoot := preorderLeft
	// 在中序遍历里找到根节点
	inorderRoot := indexMap[preorder[preOrderRoot]]

	root := &TreeNode{
		Val: preorder[preOrderRoot],
	}

	// 左子树节点数目
	leftSubTreeNums := inorderRoot - inorderLeft

	root.Left = myBuildTree(preorder, inorder, indexMap, preorderLeft+1, preorderLeft+leftSubTreeNums, inorderLeft, inorderRoot-1)
	root.Right = myBuildTree(preorder, inorder, indexMap, preorderLeft+leftSubTreeNums+1, preorderRight, inorderRoot+1, inorderRight)
	return root
}
