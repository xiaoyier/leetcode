package bst

import "leetcode/tree/binary"

// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

// 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/


func insertIntoBST(root *binary.TreeNode, val int) *binary.TreeNode {

	if root == nil {
		return &binary.TreeNode{Val: val}
	}

	origin := root
	var parent *binary.TreeNode
	for root != nil {
		if root.Val == val {
			return origin
		} else if root.Val > val {
			parent = root
			root = root.Left
		} else {
			parent = root
			root = root.Right
		}
	}


	if parent != nil {
		if val > parent.Val {
			parent.Right = &binary.TreeNode{Val: val}
		} else {
			parent.Left = &binary.TreeNode{Val: val}
		}
	}
	return origin
}