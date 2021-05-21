package bst

import (
	"leetcode/tree/binary/bt"
)

// 给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/

// 条件: 树中至少有 2 个节点。
func getMinimumDifference(root *bt.TreeNode) int {

	// 中序遍历, 迭代
	stack := []*bt.TreeNode{}
	var diff = 1 << 63 -1
	var preNum = -1
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if preNum != -1 && root.Val - preNum < diff {
			diff = root.Val - preNum
		}
		preNum = root.Val
		root = root.Right
	}
	return diff
}



// 中序遍历, 递归
func getMinimumDifference1(root *bt.TreeNode) int {

	var diff = 1 << 63 -1
	var preNum = -1
	var dfs func(node *bt.TreeNode)
	dfs = func(node *bt.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		current := node.Val - preNum
		if current < diff && preNum != -1 {
			diff = current
		}
		preNum = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return diff
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}