package bt

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/


// 迭代
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		}

		if len(stack) == 0 {
			return result
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
	}
	return result
}


// 递归
func preorderTraversal1(root *TreeNode) []int {
	var vals = []int{}
	var dfs func(*TreeNode)
	dfs = func (node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return vals
}