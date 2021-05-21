package bt

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	result := [][]int{}
	level := []int{}
	size := 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}

		level = append(level, node.Val)
		size--
		if size == 0 {
			result = append(result, level)
			level = []int{}
			size = len(q)
		}
	}
	return result
}
