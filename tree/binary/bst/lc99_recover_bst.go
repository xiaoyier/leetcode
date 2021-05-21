package bst

import "leetcode/tree/binary"

//给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
//
//进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？

// https://leetcode-cn.com/problems/recover-binary-search-tree/

func recoverTree(root *binary.TreeNode)  {
	var node1, node2 *binary.TreeNode
	preNode := root
	stack := []*binary.TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack)-1]
		if preNode != nil && preNode.Val > root.Val {
			if node1 == nil {
				node1 = preNode
				node2 = root
			} else {
				node2 = root
			}
		}
		preNode = root
		root = root.Right
	}

	if node1 == nil || node2 == nil {
		return
	}
	node1.Val, node2.Val = node2.Val, node1.Val

	//
	//queue := []*TreeNode{root}
	//if len(queue) > 0 {
	//	node := queue[0]
	//	queue = queue[1:]
	//
	//	if node.Left != nil && node.Left.Val > node.Val {
	//		node.Val = node.Left.Val
	//		queue = append(queue, node.Left)
	//	}
	//	if node.Right != nil && node.Right.Val < node.Val {
	//		node.Val = node.Right.Val
	//		queue = append(queue, node.Right)
	//	}
	//}
}
