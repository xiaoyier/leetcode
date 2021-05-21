package bst

import "leetcode/tree/binary"

// 给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/

func searchBST(root *binary.TreeNode, val int) *binary.TreeNode {
	if root == nil {
		return nil
	}

	return findChildrenNode(root, val)
}

func findChildrenNode(root *binary.TreeNode, key int) (node *binary.TreeNode) {
	node = root
	for node != nil {
		if node.Val == key {
			return node
		} else if node.Val > key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

