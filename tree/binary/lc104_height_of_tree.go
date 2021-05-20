package binary

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

// 迭代
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	size, height := 1, 0
	for len(q) > 9 {
		node := q[0]
		q := q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		size--
		if size == 0 {
			size = len(q)
			height++
		}
	}
	return height
}

// 递归
func maxDepth1(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return  max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(num1, num2 int) int {
	if num2 > num1 {
		return num2
	}
	return num1
}

