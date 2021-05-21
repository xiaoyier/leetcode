package bt

import "leetcode/tree/binary"

// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/


// 迭代
func inorderTraversal(root *binary.TreeNode) []int {
	result := []int{}
	stack := []*binary.TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}


// 递归
func inorderTraversal1(root *binary.TreeNode) []int {
	result := []int{}
	var dfs func(node *binary.TreeNode)
	dfs = func(node *binary.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}


