package binary

// 给定一个二叉树，返回它的 后序 遍历。

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// 迭代
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	var preNode *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == preNode {
			result = append(result, root.Val)
			preNode = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return result
}

// 递归
func postorderTraversal1(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}
