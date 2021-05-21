package bst

import "leetcode/tree/binary"

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

// 一般来说，删除节点可分为两个步骤：
//
// 首先找到需要删除的节点；
// 如果找到了，删除它。
// 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

// https://leetcode-cn.com/problems/delete-node-in-a-bst/

func deleteNode(root *binary.TreeNode, key int) *binary.TreeNode {

	deleteNode, parent := findNode(root, key)
	if deleteNode == nil {
		return root
	}
	// 双度节点
	if deleteNode.Left != nil && deleteNode.Right != nil {
		successorNode, p := successor(deleteNode)
		deleteNode.Val = successorNode.Val
		deleteNode = successorNode
		parent = p
	}

	child := deleteNode.Left
	if deleteNode.Right != nil {
		child = deleteNode.Right
	}


	if deleteNode == root {
		root = child
	} else if deleteNode == parent.Left {
		parent.Left = child
	} else  {
		parent.Right = child
	}

	return root
}

// O(logn) 即 O(h)
func findNode(root *binary.TreeNode, key int) (node *binary.TreeNode, parent *binary.TreeNode) {
	node = root
	for node != nil {
		if node.Val == key {
			return node, parent
		} else if node.Val > key {
			parent = node
			node = node.Left
		} else {
			parent = node
			node = node.Right
		}
	}
	return node, parent
}


func successor(node *binary.TreeNode) (*binary.TreeNode, *binary.TreeNode) {
	if node == nil || node.Right == nil {
		return nil, nil
	}

	parent := node
	node = node.Right
	for node.Left != nil {
		parent = node
		node = node.Left
	}
	return node, parent
}