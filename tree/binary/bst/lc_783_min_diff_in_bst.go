package bst

import "leetcode/tree/binary"

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/

func minDiffInBST(root *binary.TreeNode) int {
	diff := 1 << 63 - 1
	pre := -1
	stack := []*binary.TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if pre != -1 && diff > root.Val - pre {
			diff = root.Val - pre
		}
		pre = root.Val
		root = root.Right
	}

	return diff
}