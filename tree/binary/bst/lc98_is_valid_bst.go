package bst

import (
	"leetcode/tree/binary/bt"
)

// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 假设一个二叉搜索树具有如下特征：
//
// 	节点的左子树只包含小于当前节点的数。
// 	节点的右子树只包含大于当前节点的数。
// 	所有左子树和右子树自身必须也是二叉搜索树。

// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST1(root *bt.TreeNode) bool {

	stack := []*bt.TreeNode{}
	inOrder := -1 << 63

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inOrder {
			return false
		}
		inOrder = root.Val
		root = root.Right
	}

	return false
}

var pre = -1 << 63
func isValidBST(root *bt.TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST(root.Right)
}

func isValidBST2(root *bt.TreeNode) bool {

	nodes := make([]*bt.TreeNode, 0)
	inOrder(root, func(node *bt.TreeNode) {
		nodes = append(nodes, node)
	})

	if len(nodes) < 2 {
		return true
	}
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Val <= nodes[i - 1].Val {
			return false
		}
	}
	return true
}

func inOrder(root *bt.TreeNode, it func(node *bt.TreeNode)) {
	if root == nil {
		return
	}

	inOrder(root.Left, it)
	it(root)
	inOrder(root.Right, it)
}