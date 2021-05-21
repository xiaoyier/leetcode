package bt

import "leetcode/tree/binary"

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

func levelOrderBottom(root *binary.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*binary.TreeNode{root}
	stack := [][]int{}
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
			stack = append(stack, level)
			level = []int{}
			size = len(q)
		}
	}

	result := make([][]int, len(stack))
	for i := 0; i < len(stack) ; i++ {
		result[i] = stack[len(stack) -1 - i]
	}
	return result
}
