package bt

// 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

// 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	root.Val = 1
	q := []*TreeNode{root}
	size, width := 1, 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.Left != nil {
			node.Left.Val = node.Val * 2
			q = append(q, node.Left)
		}

		if node.Right != nil {
			node.Right.Val = node.Val * 2 + 1
			q= append(q, node.Right)
		}

		size--
		if size == 0 && len(q) > 0 {
			currW := q[len(q)-1].Val - q[0].Val + 1
			if currW > width {
				width = currW
			}
			size = len(q)
		}
	}
	return width
}
