package binary


// 给定一个 N 叉树，返回其节点值的 前序遍历 。
//
//N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {

	vals := []int{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		for _, item := range node.Children {
			dfs(item)
		}
	}
	dfs(root)
	return vals

}