package bst

import "leetcode/tree/binary"

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

func kthSmallest(root *binary.TreeNode, k int) int {
	index := 0
	result := -1
	var dfs func(node *binary.TreeNode)
	dfs = func(node *binary.TreeNode) {
		if node == nil || result >= 0 {
			return
		}

		dfs(node.Left)
		index++
		if index == k {
			result = node.Val
			return
		}
		dfs(node.Right)
	}

	dfs(root)
	return result
}
