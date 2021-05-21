package bst

import "leetcode/tree/binary"

// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

// https://leetcode-cn.com/problems/range-sum-of-bst/

func rangeSumBST(root *binary.TreeNode, low int, high int) int {
	queue := []*binary.TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			queue = append(queue, node.Right)
		} else if node.Val > high {
			queue = append(queue, node.Left)
		} else {
			sum += node.Val
			queue = append(queue, node.Left, node.Right)
		}
	}
	return sum
}

func rangeSumBST1(root *binary.TreeNode, low int, high int) int {
	sum := 0
	var rangeFunc func(node *binary.TreeNode)
	rangeFunc = func(node *binary.TreeNode) {
		if node == nil {
			return
		}

		rangeFunc(node.Left)
		if node.Val >= low && node.Val <= high {
			sum += node.Val
		}
		rangeFunc(node.Right)
	}
	rangeFunc(root)
	return sum
}

