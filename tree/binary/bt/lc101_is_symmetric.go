package bt

// 给定一个二叉树，检查它是否是镜像对称的。

// https://leetcode-cn.com/problems/symmetric-tree/

// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3

// 中序遍历
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func (node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		return  node1.Val == node2.Val && dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root.Left, root.Right)
}
