package bst

import (
	"leetcode/tree/binary/bt"
)

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *bt.TreeNode) *bt.TreeNode {
	node := root
	for node != nil {
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		} else {
			return node
		}
	}
	return node
}
